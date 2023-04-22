package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DoraTiger/NEU_IPGW/pkg/data"
)

func GetReadableLoginMessage(resp *http.Response) (int, string, error) {
	var message data.LoginMessage
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &message)
	if err != nil {
		return 1, "", fmt.Errorf("Error format message: %v", err)
	}
	return message.Code, (message.Message), nil
}

func GetCNLoginMessage(m string) string {
	switch m {
	case "ip_already_online_error":
		return "IP已经在线"
	case "success":
		return "登录成功"
	default:
		return "未知错误"

	}
}
