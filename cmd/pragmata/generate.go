package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gear6io/pragmata/pkg/config"
	httpserver "github.com/gear6io/pragmata/pkg/http/server"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/modules/sources"
	"github.com/gear6io/pragmata/pkg/modules/suggestions"
	"github.com/gear6io/pragmata/pkg/sqlstore"
)

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate artifacts",
	}
	cmd.AddCommand(generateOpenAPICmd())
	return cmd
}

func generateOpenAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "openapi",
		Short: "Write OpenAPI spec to docs/api/openapi.yaml",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runGenerateOpenAPI()
		},
	}
}

func runGenerateOpenAPI() error {
	// Address is irrelevant for spec generation; use defaults if config is absent.
	cfg, err := config.Load(configPath)
	if err != nil {
		cfg = config.Defaults()
	}

	srv := httpserver.New(
		httpserver.NewProvider(
			struct{ pipes.Handler }{},
			struct{ suggestions.Handler }{},
			struct{ sources.Handler }{},
		),
		struct{ sqlstore.SQLStore }{},
		httpserver.Addr(cfg.Server.Host, cfg.Server.Port),
	)

	data, err := srv.OpenAPISpec()
	if err != nil {
		return fmt.Errorf("marshal spec: %w", err)
	}

	const outPath = "docs/api/openapi.yaml"
	if err := os.MkdirAll("docs/api", 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}
	if err := os.WriteFile(outPath, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}
	fmt.Printf("wrote %s\n", outPath)
	return nil
}
