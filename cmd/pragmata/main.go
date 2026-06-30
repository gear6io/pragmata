package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/gear6io/pragmata/pkg/config"
	"github.com/gear6io/pragmata/pkg/datastore"
	errors "github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/executor/goroutineexecutor"
	httpserver "github.com/gear6io/pragmata/pkg/http/server"
	"github.com/gear6io/pragmata/pkg/modules/pipes/implpipes"
	"github.com/gear6io/pragmata/pkg/querier"
	"github.com/gear6io/pragmata/pkg/modules/sources/implsources"
	"github.com/gear6io/pragmata/pkg/modules/suggestions/implsuggestions"
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
	root.AddCommand(generateCmd())
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
		return errors.WrapInternalf(err, errors.CodeInternal, "load config")
	}

	// Storage
	store, err := sqlitesqlstore.New(cfg.Database.Path)
	if err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "open database")
	}

	// Migrations
	migrations, err := sqlmigration.New([]sqlmigration.SQLMigration{
		sqlmigration.NewInitialSchema(),
	})
	if err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "build migrations")
	}
	if err := sqlmigrator.New(store.BunDB(), migrations).Migrate(ctx); err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "migrate")
	}

	// SQLMesh runner
	runner := sqlmesh.New(cfg.SQLMesh.ProjectDir, cfg.SQLMesh.BinaryPath)
	if err := runner.EnsureProject(cfg.ClickHouse.URL); err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "bootstrap sqlmesh project")
	}

	// Executor
	exec := goroutineexecutor.New(store, runner)
	if err := exec.ResumeInterrupted(ctx); err != nil {
		log.Printf("warn: resume interrupted jobs: %v", err)
	}

	// Scheduler — loads its own pipes from the store on Start
	scheduler := cronscheduler.New(exec, store)
	if err := scheduler.Start(ctx); err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "start scheduler")
	}
	defer func() { _ = scheduler.Stop(context.Background()) }()

	// DataStore (ClickHouse) + sources module + handler
	ds, err := datastore.New(cfg.ClickHouse)
	if err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "open datastore")
	}

	// Querier (dedicated ClickHouse connection for ENDPOINT pipe execution)
	q, err := querier.New(cfg.ClickHouse.URL)
	if err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "open querier")
	}

	// Pipes module + handler
	mod := implpipes.NewModule(ds, store, exec, scheduler, q)
	h := implpipes.NewHandler(mod)

	// Suggestions module + handler
	suggestMod := implsuggestions.NewModule(store)
	suggestH := implsuggestions.NewHandler(suggestMod)

	sourceMod := implsources.NewModule(ds)
	sourceH := implsources.NewHandler(sourceMod)

	// HTTP server
	addr := httpserver.Addr(cfg.Server.Host, cfg.Server.Port)
	srv := httpserver.New(httpserver.NewProvider(h, suggestH, sourceH), store, addr)
	log.Printf("pragmata listening on %s", addr)

	// Graceful shutdown on SIGINT/SIGTERM
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	return srv.Start(ctx)
}
