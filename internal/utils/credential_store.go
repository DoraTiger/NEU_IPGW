package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/DoraTiger/NEU_IPGW/config"
)

var ErrAccountNotFound = errors.New("saved account not found")

type SavedCredential struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

type credentialContainer struct {
	Version     int                         `json:"version"`
	LastAccount string                      `json:"last_account"`
	Accounts    map[string]*SavedCredential `json:"accounts"`
}

type AccountMeta struct {
	Username  string
	UpdatedAt time.Time
	IsLast    bool
}

type CredentialStore struct {
	configDir string
	filePath  string
	masterKey string
}

func NewCredentialStore(configDir string, masterKey string) (*CredentialStore, error) {
	if err := EnsureConfigDir(configDir); err != nil {
		return nil, err
	}
	if masterKey == "" {
		return nil, errors.New("master key is empty")
	}

	return &CredentialStore{
		configDir: configDir,
		filePath:  filepath.Join(configDir, config.CredentialFileName),
		masterKey: masterKey,
	}, nil
}

func (s *CredentialStore) Save(username string, password string) error {
	container, err := s.loadContainer()
	if err != nil {
		return err
	}
	if container.Accounts == nil {
		container.Accounts = map[string]*SavedCredential{}
	}

	now := time.Now().UTC()
	container.Accounts[username] = &SavedCredential{
		Username:  username,
		Password:  password,
		UpdatedAt: now,
	}
	container.LastAccount = username
	return s.saveContainer(container)
}

func (s *CredentialStore) LoadByUsername(username string) (*SavedCredential, error) {
	container, err := s.loadContainer()
	if err != nil {
		return nil, err
	}

	cred, ok := container.Accounts[username]
	if !ok {
		return nil, ErrAccountNotFound
	}
	return cred, nil
}

func (s *CredentialStore) LoadLast() (*SavedCredential, error) {
	container, err := s.loadContainer()
	if err != nil {
		return nil, err
	}
	if container.LastAccount == "" {
		return nil, ErrAccountNotFound
	}
	cred, ok := container.Accounts[container.LastAccount]
	if !ok {
		return nil, ErrAccountNotFound
	}
	return cred, nil
}

func (s *CredentialStore) DeleteByUsername(username string) (bool, error) {
	container, err := s.loadContainer()
	if err != nil {
		return false, err
	}
	if _, ok := container.Accounts[username]; !ok {
		return false, nil
	}

	delete(container.Accounts, username)
	if container.LastAccount == username {
		container.LastAccount = ""
		for name := range container.Accounts {
			container.LastAccount = name
			break
		}
	}

	if len(container.Accounts) == 0 {
		if err := os.Remove(s.filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
			return false, fmt.Errorf("delete credential file failed: %w", err)
		}
		return true, nil
	}

	return true, s.saveContainer(container)
}

func (s *CredentialStore) DeleteLast() (string, bool, error) {
	container, err := s.loadContainer()
	if err != nil {
		return "", false, err
	}
	if container.LastAccount == "" {
		return "", false, nil
	}

	username := container.LastAccount
	ok, err := s.DeleteByUsername(username)
	return username, ok, err
}

func (s *CredentialStore) MarkLast(username string) error {
	container, err := s.loadContainer()
	if err != nil {
		return err
	}
	if _, ok := container.Accounts[username]; !ok {
		return ErrAccountNotFound
	}
	container.LastAccount = username
	return s.saveContainer(container)
}

func (s *CredentialStore) ListAccounts() ([]AccountMeta, error) {
	container, err := s.loadContainer()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []AccountMeta{}, nil
		}
		return nil, err
	}

	metas := make([]AccountMeta, 0, len(container.Accounts))
	for username, cred := range container.Accounts {
		metas = append(metas, AccountMeta{
			Username:  username,
			UpdatedAt: cred.UpdatedAt,
			IsLast:    username == container.LastAccount,
		})
	}
	// Keep stable output for CLI.
	sort.Slice(metas, func(i, j int) bool {
		if metas[i].UpdatedAt.Equal(metas[j].UpdatedAt) {
			return metas[i].Username < metas[j].Username
		}
		return metas[i].UpdatedAt.After(metas[j].UpdatedAt)
	})
	return metas, nil
}

func (s *CredentialStore) loadContainer() (*credentialContainer, error) {
	buf, err := os.ReadFile(s.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &credentialContainer{Version: 1, Accounts: map[string]*SavedCredential{}}, nil
		}
		return nil, fmt.Errorf("read credential file failed: %w", err)
	}

	plain, err := DecryptWithMasterKey(buf, s.masterKey)
	if err != nil {
		return nil, fmt.Errorf("decrypt credential file failed: %w", err)
	}

	container := &credentialContainer{}
	if err := json.Unmarshal(plain, container); err != nil {
		return nil, fmt.Errorf("parse credential file failed: %w", err)
	}
	if container.Accounts == nil {
		container.Accounts = map[string]*SavedCredential{}
	}
	return container, nil
}

func (s *CredentialStore) saveContainer(container *credentialContainer) error {
	container.Version = 1
	raw, err := json.Marshal(container)
	if err != nil {
		return fmt.Errorf("marshal credential data failed: %w", err)
	}
	cipherBuf, err := EncryptWithMasterKey(raw, s.masterKey)
	if err != nil {
		return err
	}

	tmpPath := s.filePath + ".tmp"
	if err := os.WriteFile(tmpPath, cipherBuf, 0o600); err != nil {
		return fmt.Errorf("write temp credential file failed: %w", err)
	}
	if err := os.Rename(tmpPath, s.filePath); err != nil {
		return fmt.Errorf("replace credential file failed: %w", err)
	}
	if err := os.Chmod(s.filePath, 0o600); err != nil {
		return fmt.Errorf("set credential file permission failed: %w", err)
	}
	return nil
}
