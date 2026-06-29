package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/gear6io/pragmata/pkg/errors"
)

type Config struct {
	Server     ServerConfig     `yaml:"server"`
	ClickHouse ClickHouseConfig `yaml:"clickhouse"`
	SQLMesh    SQLMeshConfig    `yaml:"sqlmesh"`
	Database   DatabaseConfig   `yaml:"database"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type ClickHouseConfig struct {
	URL string `yaml:"url"`
}

type SQLMeshConfig struct {
	ProjectDir string `yaml:"project_dir"`
	BinaryPath string `yaml:"binary_path"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.WrapInternalf(err, errors.CodeInternal, "read config %q", path)
	}
	cfg := defaults()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, errors.WrapInvalidInputf(err, errors.CodeInvalidInput, "parse config")
	}
	return cfg, nil
}

// Defaults returns the default configuration without reading any file.
func Defaults() *Config { return defaults() }

func defaults() *Config {
	return &Config{
		Server:     ServerConfig{Host: "0.0.0.0", Port: 7181},
		ClickHouse: ClickHouseConfig{URL: "clickhouse://localhost:9000"},
		SQLMesh:    SQLMeshConfig{BinaryPath: "sqlmesh", ProjectDir: ".pragmata/sqlmesh"},
		Database:   DatabaseConfig{Path: "./pragmata.db"},
	}
}
