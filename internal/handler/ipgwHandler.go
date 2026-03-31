package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/DoraTiger/NEU_IPGW/config"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/sirupsen/logrus"
)

type ipgwHandler struct {
	client *http.Client
	logger *logrus.Logger
}

var ErrIPGWNotOnline = errors.New("ipgw not online")

type OnlineInfo struct {
	Username    string
	UsedBytes   float64
	UsedSeconds int64
	Balance     float64
	IP          string
}

func (o OnlineInfo) Traffic() string {
	return formatBytes(o.UsedBytes)
}

func (o OnlineInfo) Duration() string {
	return formatDuration(o.UsedSeconds)
}

type radUserInfoResponse struct {
	Error       string      `json:"error"`
	UserName    string      `json:"user_name"`
	SumBytes    json.Number `json:"sum_bytes"`
	SumSeconds  json.Number `json:"sum_seconds"`
	UserBalance json.Number `json:"user_balance"`
	OnlineIP    string      `json:"online_ip"`
}

func extractJSONFromJSONP(body []byte) ([]byte, error) {
	payload := strings.TrimSpace(string(body))
	start := strings.Index(payload, "(")
	end := strings.LastIndex(payload, ")")
	if start == -1 || end == -1 || end <= start {
		return nil, fmt.Errorf("unexpected JSONP response: %s", payload)
	}
	return []byte(payload[start+1 : end]), nil
}

func (h *ipgwHandler) FetchOnlineInfo() (*OnlineInfo, error) {
	statusURL, err := url.Parse(config.DefaultIPGatewayStatusURL)
	if err != nil {
		return nil, err
	}
	query := statusURL.Query()
	if query.Get("callback") == "" {
		query.Set("callback", "json")
	}
	statusURL.RawQuery = query.Encode()

	req, _ := http.NewRequest("GET", statusURL.String(), nil)
	req.Header.Set("Referer", config.DefaultIPGatewayLoggedinURL)
	resp, err := h.client.Do(req)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return nil, err
	}

	jsonBody, err := extractJSONFromJSONP(body)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return nil, err
	}

	var payload radUserInfoResponse
	if err := json.Unmarshal(jsonBody, &payload); err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return nil, err
	}

	if payload.Error != "ok" {
		return nil, interpretStatusError(payload.Error)
	}

	usedBytes, err := parseFloat(payload.SumBytes)
	if err != nil {
		return nil, err
	}
	usedSeconds, err := parseInt(payload.SumSeconds)
	if err != nil {
		return nil, err
	}
	balance, err := parseFloat(payload.UserBalance)
	if err != nil {
		return nil, err
	}

	return &OnlineInfo{
		Username:    payload.UserName,
		UsedBytes:   usedBytes,
		UsedSeconds: usedSeconds,
		Balance:     balance,
		IP:          payload.OnlineIP,
	}, nil
}

func formatBytes(value float64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	idx := 0
	for value >= 1024 && idx < len(units)-1 {
		value /= 1024
		idx++
	}
	return fmt.Sprintf("%.2f %s", value, units[idx])
}

func formatDuration(seconds int64) string {
	if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	}
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	if minutes < 60 {
		return fmt.Sprintf("%dm %ds", minutes, remainingSeconds)
	}
	hours := minutes / 60
	remainingMinutes := minutes % 60
	if hours < 24 {
		return fmt.Sprintf("%dh %dm", hours, remainingMinutes)
	}
	days := hours / 24
	remainingHours := hours % 24
	return fmt.Sprintf("%dd %dh", days, remainingHours)
}

func parseFloat(value json.Number) (float64, error) {
	if strings.TrimSpace(value.String()) == "" {
		return 0, nil
	}
	return value.Float64()
}

func parseInt(value json.Number) (int64, error) {
	if strings.TrimSpace(value.String()) == "" {
		return 0, nil
	}
	return value.Int64()
}

func interpretStatusError(code string) error {
	switch code {
	case "not_online_error", "logout_error", "ip_not_online_error", "logout", "not_online":
		return ErrIPGWNotOnline
	default:
		return fmt.Errorf("ipgw status error: %s", code)
	}
}

// set logger
func (h *ipgwHandler) SetLogger(logger *logrus.Logger) {
	h.logger = logger
}

func NewIPGWHandler() *ipgwHandler {
	return &ipgwHandler{
		client: utils.NewClientWithJar(),
	}
}

func (h *ipgwHandler) SetClient(client *http.Client) {
	h.client = client
}

// login to ipgw
// https://ipgw.neu.edu.cn/
func (h *ipgwHandler) Login() (string, error) {
	// get query parameters corresponding to the gateway url on the current network
	resp, err := h.client.Get(config.DefaultIPGatewayLoginURL)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return "", err
	}
	defer resp.Body.Close()

	// get ticket
	resp, err = h.client.Get(config.DefaultIPGatewayTicketURL + resp.Request.URL.RawQuery)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return "", err
	}
	defer resp.Body.Close()

	// login by ticket
	req, _ := http.NewRequest("GET", config.DefaultIPGatewayAPIURL+resp.Request.URL.RequestURI(), nil)
	resp, err = h.client.Do(req)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return "", err
	}
	defer resp.Body.Close()

	//get final result
	_, msg, err := utils.GetReadableLoginMessage(resp)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return "", err
	}
	return msg, nil

}

// logout from ipgw
// https://ipgw.neu.edu.cn/
func (h *ipgwHandler) Logout() error {
	req, _ := http.NewRequest("GET", config.DefaultIPGatewayLogoutURL, nil)
	req.Header.Add("Referer", config.DefaultIPGatewayLoggedinURL)
	_, err := h.client.Do(req)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return err
	}
	return nil
}
