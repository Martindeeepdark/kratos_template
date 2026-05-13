package conf

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Bootstrap struct {
	Server ServerConfig `yaml:"server"`
	Data   DataConfig   `yaml:"data"`
}

type ServerConfig struct {
	HTTP HTTPConfig `yaml:"http"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type HTTPConfig struct {
	Network string        `yaml:"network"`
	Addr    string        `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
}

type GRPCConfig struct {
	Network string        `yaml:"network"`
	Addr    string        `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
}

type DataConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

func Load(path string) (*Bootstrap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var bc Bootstrap
	if err := yaml.Unmarshal(data, &bc); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return &bc, nil
}
