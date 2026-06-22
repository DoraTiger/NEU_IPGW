package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type SSOHandler struct {
	client *http.Client
	logger *logrus.Logger
}

func NewSSOHandler(client *http.Client) *SSOHandler {
	return &SSOHandler{
		client: client,
	}
}

func (h *SSOHandler) SetLogger(logger *logrus.Logger) {
	h.logger = logger
}

func (h *SSOHandler) Login(casLoginURL string) error {
	resp, err := h.client.Get(casLoginURL)
	if err != nil {
		return fmt.Errorf("CAS login request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("CAS login HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (h *SSOHandler) Logout(logoutURL string) error {
	req, _ := http.NewRequest("GET", logoutURL, nil)
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
