package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

func Load() (*Config, error) {
	path, err := EnsureConfig()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}