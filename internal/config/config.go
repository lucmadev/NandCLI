package config

import (
	_ "embed"
	"os"
	"path/filepath"
)

var defaultConfig []byte

func ConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "nandcli"), nil
}

func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "config.yaml"), nil
}

func EnsureConfig() (string, error) {
	path, err := ConfigPath()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return "", err
		}

		if err := os.WriteFile(path, defaultConfig, 0644); err != nil {
			return "", err
		}
	}

	return path, nil
}