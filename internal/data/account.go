package data

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/DoraTiger/NEU_IPGW/config"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAccount(username string, password string) *Account {
	return &Account{
		Username: username,
		Password: password,
	}
}

func (a *Account) GetUserNameLength() int {
	return len(a.Username)
}

func (a *Account) GetPasswordLength() int {
	return len(a.Password)
}

func (a *Account) GetRSA(ltID string) string {
	plainText := fmt.Sprintf("%s%s", a.Username, a.Password)
	// undefine
	pubKey, err := parseRSAPublicKey(config.DefaultPublicKeyStr)
	if err != nil {
		panic(fmt.Sprintf("failed to parse public key: %v", err))
	}
	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(plainText))
	if err != nil {
		panic(fmt.Sprintf("failed to encrypt: %v", err))
	}
	return base64.StdEncoding.EncodeToString(cipherBytes)
}

// === 私有工具函数 ===

// parseRSAPublicKey 从 Base64 字符串解析 PKCS#1 或 PKIX 格式的 RSA 公钥
func parseRSAPublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	der, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}

	// 尝试 PKIX (SubjectPublicKeyInfo) —— 适用于标准 PEM 公钥（带头尾的那种）
	if pub, err := x509.ParsePKIXPublicKey(der); err == nil {
		if rsaPub, ok := pub.(*rsa.PublicKey); ok {
			return rsaPub, nil
		}
	}

	// 尝试 PKCS#1（裸 RSA 公钥，无 ASN.1 wrapper）—— 你的公钥属于这种
	if rsaPub, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return rsaPub, nil
	}

	return nil, fmt.Errorf("failed to parse public key as either PKIX or PKCS#1")
}
