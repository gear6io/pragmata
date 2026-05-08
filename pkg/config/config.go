package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}
	cfg := defaults()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	return cfg, nil
}

func defaults() *Config {
	return &Config{
		Server:     ServerConfig{Host: "0.0.0.0", Port: 7181},
		ClickHouse: ClickHouseConfig{Host: "localhost", Port: 9000, Database: "default"},
		SQLMesh:    SQLMeshConfig{BinaryPath: "sqlmesh", ProjectDir: ".pragmata/sqlmesh"},
		Database:   DatabaseConfig{Path: "./pragmata.db"},
	}
}
