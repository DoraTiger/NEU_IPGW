package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DoraTiger/NEU_IPGW/config"
)

func ResolveConfigDir(cliValue string) (string, error) {
	if envValue, ok := os.LookupEnv(config.EnvConfigDir); ok {
		envValue = strings.TrimSpace(envValue)
		if envValue != "" {
			return envValue, nil
		}
	}

	cliValue = strings.TrimSpace(cliValue)
	if cliValue != "" {
		return cliValue, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir failed: %w", err)
	}
	return filepath.Join(homeDir, config.DefaultConfigSubPath), nil
}

func ResolveMasterKey(cliValue string) string {
	if envValue, ok := os.LookupEnv(config.EnvMasterKey); ok {
		envValue = strings.TrimSpace(envValue)
		if envValue != "" {
			return envValue
		}
	}

	cliValue = strings.TrimSpace(cliValue)
	if cliValue != "" {
		return cliValue
	}

	return config.DefaultMasterKey
}

func EnsureConfigDir(dir string) error {
	if strings.TrimSpace(dir) == "" {
		return errors.New("config dir is empty")
	}

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create config dir failed: %w", err)
	}
	if err := os.Chmod(dir, 0o700); err != nil {
		return fmt.Errorf("set config dir permission failed: %w", err)
	}
	return nil
}
