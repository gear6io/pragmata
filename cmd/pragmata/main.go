package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/gear6io/pragmata/pkg/config"
	httpserver "github.com/gear6io/pragmata/pkg/http/server"
	"github.com/gear6io/pragmata/pkg/modules/pipes/implpipes"
	"github.com/gear6io/pragmata/pkg/orchestration/goroutineorchestration"
	"github.com/gear6io/pragmata/pkg/scheduler/cronscheduler"
	"github.com/gear6io/pragmata/pkg/sqlmesh"
	"github.com/gear6io/pragmata/pkg/sqlmigration"
	"github.com/gear6io/pragmata/pkg/sqlmigrator"
	"github.com/gear6io/pragmata/pkg/sqlstore/sqlitesqlstore"
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
	store, err := sqlitesqlstore.New(cfg.Database.Path)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}

	// Migrations
	migrations, err := sqlmigration.New([]sqlmigration.SQLMigration{
		sqlmigration.NewInitialSchema(),
	})
	if err != nil {
		return fmt.Errorf("build migrations: %w", err)
	}
	if err := sqlmigrator.New(store.BunDB(), migrations).Migrate(ctx); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	// SQLMesh runner (Phase B stubs)
	runner := sqlmesh.New(cfg.SQLMesh.ProjectDir, cfg.SQLMesh.BinaryPath)

	// Orchestrator
	orchest := goroutineorchestration.New(store, runner)
	if err := orchest.ResumeInterrupted(ctx); err != nil {
		log.Printf("warn: resume interrupted jobs: %v", err)
	}

	// Scheduler — loads its own pipes from the store on Start
	sched := cronscheduler.New(orchest, store)
	if err := sched.Start(ctx); err != nil {
		return fmt.Errorf("start scheduler: %w", err)
	}
	defer func() { _ = sched.Stop(context.Background()) }()

	// Module + handler
	mod := implpipes.NewModule(store, orchest, sched)
	h := implpipes.NewHandler(mod)

	// HTTP server
	addr := httpserver.Addr(cfg.Server.Host, cfg.Server.Port)
	srv := httpserver.New(h, store, addr)
	log.Printf("pragmata listening on %s", addr)

	// Graceful shutdown on SIGINT/SIGTERM
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	return srv.Start(ctx)
}
