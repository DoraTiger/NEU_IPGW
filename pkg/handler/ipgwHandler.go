package handler

import (
	"fmt"
	"net/http"

	"github.com/DoraTiger/NEU_IPGW/config"
	"github.com/DoraTiger/NEU_IPGW/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ipgwHandler struct {
	client *http.Client
	logger *logrus.Logger
}

//set logger
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
