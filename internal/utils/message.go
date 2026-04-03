package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginMessage struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Redirect string `json:"Redirect"`
	ID       string `json:"ID"`
}

func GetReadableLoginMessage(resp *http.Response) (int, string, error) {
	var message LoginMessage
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &message)
	if err != nil {
		return 1, "", fmt.Errorf("Error format message: %v", err)
	}
	return message.Code, (message.Message), nil
}

func GetCNLoginMessage(m string) string {
	switch m {
	case "ip_already_online_error":
		return "IP 已经在线"
	case "success":
		return "登录成功"
	default:
		return "未知错误"

	}
}
