package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/gear6io/pragmata/internal/config"
	"github.com/gear6io/pragmata/internal/executor"
	gorchest "github.com/gear6io/pragmata/internal/orchestration/goroutine"
	"github.com/gear6io/pragmata/internal/scheduler"
	"github.com/gear6io/pragmata/internal/server"
	"github.com/gear6io/pragmata/internal/sqlmesh"
	internalsqlstore "github.com/gear6io/pragmata/internal/sqlstore"
	"github.com/gear6io/pragmata/pkg/modules/pipes/implpipes"
)

var configPath string

func main() {
	root := &cobra.Command{
		Use:   "pragmata",
		Short: "ClickHouse developer platform — open-source TinyBird alternative",
	}
	root.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yaml", "path to config.yaml")
	root.AddCommand(serveCmd())
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP API server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return serve(cmd.Context())
		},
	}
}

func serve(ctx context.Context) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	// Storage
	store, err := internalsqlstore.New(cfg.Database.Path)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}

	// ClickHouse executor
	exec, err := executor.NewClickHouseExecutor(cfg.ClickHouse)
	if err != nil {
		return fmt.Errorf("clickhouse: %w", err)
	}
	defer exec.Close()

	// SQLMesh runner (Phase B stubs)
	runner := sqlmesh.New(cfg.SQLMesh.ProjectDir, cfg.SQLMesh.BinaryPath)

	// Orchestrator
	orchest := gorchest.New(store, runner)
	if err := orchest.ResumeInterrupted(ctx); err != nil {
		log.Printf("warn: resume interrupted jobs: %v", err)
	}

	// Scheduler
	allPipes, err := store.ListPipes(ctx)
	if err != nil {
		return fmt.Errorf("list pipes: %w", err)
	}
	sched := scheduler.New(orchest)
	if err := sched.Start(allPipes); err != nil {
		return fmt.Errorf("start scheduler: %w", err)
	}
	defer sched.Stop()

	// Module + handler
	mod := implpipes.NewModule(store, orchest, exec, sched)
	h := implpipes.NewHandler(mod)

	// HTTP server
	srv := server.New(h, store)
	addr := server.Addr(cfg.Server.Host, cfg.Server.Port)
	log.Printf("pragmata listening on %s", addr)

	// Graceful shutdown on SIGINT/SIGTERM
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	return srv.Start(ctx, addr)
}
