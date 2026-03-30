package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

func DeriveAES256Key(masterKey string) []byte {
	sum := sha256.Sum256([]byte(masterKey))
	return sum[:]
}

func EncryptWithMasterKey(plain []byte, masterKey string) ([]byte, error) {
	key := DeriveAES256Key(masterKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("init cipher failed: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("init gcm failed: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("generate nonce failed: %w", err)
	}

	cipherText := gcm.Seal(nil, nonce, plain, nil)
	payload := append(nonce, cipherText...)
	return payload, nil
}

func DecryptWithMasterKey(payload []byte, masterKey string) ([]byte, error) {
	key := DeriveAES256Key(masterKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("init cipher failed: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("init gcm failed: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(payload) < nonceSize {
		return nil, errors.New("cipher payload is too short")
	}

	nonce := payload[:nonceSize]
	cipherText := payload[nonceSize:]
	plain, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypt payload failed: %w", err)
	}
	return plain, nil
}
