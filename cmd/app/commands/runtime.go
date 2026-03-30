package commands

import (
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/DoraTiger/NEU_IPGW/config"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
)

var (
	configDirFlag string
	masterKeyFlag string
	runtimeCfgDir string
	runtimeKey    string
)

func initRuntimeConfig() error {
	resolvedDir, err := utils.ResolveConfigDir(configDirFlag)
	if err != nil {
		return err
	}
	if err := utils.EnsureConfigDir(resolvedDir); err != nil {
		return err
	}

	runtimeCfgDir = resolvedDir
	runtimeKey = utils.ResolveMasterKey(masterKeyFlag)

	warnLegacyDirIfNeeded(resolvedDir)
	return nil
}

func newCredentialStore() (*utils.CredentialStore, error) {
	if runtimeCfgDir == "" || runtimeKey == "" {
		return nil, fmt.Errorf("runtime config not initialized")
	}
	return utils.NewCredentialStore(runtimeCfgDir, runtimeKey)
}

func warnLegacyDirIfNeeded(resolvedDir string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	legacyDir := filepath.Join(homeDir, cfg.DefaultNEUDir)
	if _, err := os.Stat(legacyDir); err == nil {
		credFile := filepath.Join(resolvedDir, cfg.CredentialFileName)
		if _, err := os.Stat(credFile); os.IsNotExist(err) {
			logger.Warn("legacy directory .NEU detected; credentials now use ~/.config/doratiger/neu-ipgw")
		}
	}
}
