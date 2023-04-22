package data

import "fmt"

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
	return fmt.Sprintf("%s%s%s", a.Username, a.Password, ltID)
}
