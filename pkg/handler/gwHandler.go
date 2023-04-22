package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"errors"

	"github.com/DoraTiger/NEU_IPGW/config"
	"github.com/DoraTiger/NEU_IPGW/pkg/data"
	"github.com/DoraTiger/NEU_IPGW/pkg/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type GwHandler struct {
	client *http.Client
	logger *logrus.Logger
}

//set logger
func (h *GwHandler) SetLogger(logger *logrus.Logger) {
	h.logger = logger
}

// get client
func (h *GwHandler) GetClient() *http.Client {
	return h.client
}

func NewGWHandler() *GwHandler {
	return &GwHandler{
		client: utils.NewClientWithJar(),
	}
}

func (h *GwHandler) Login(username string, password string) error {
	account := data.NewAccount(username, password)
	err := h.LoginWithAccount(account)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return err
	}
	return nil
}

func (h *GwHandler) LoginWithAccount(account *data.Account) error {
	// get ltID, executionID
	req, _ := http.NewRequest("GET", config.DefaultGatewayURL, nil)
	resp, err := h.client.Do(req)

	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return err
	}
	defer resp.Body.Close()

	ltID, executionID, err := getLtAndExecution(resp)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return err
	}

	// get cookie
	// Parameter definition reference https://pass.neu.edu.cn/tpass/comm/neu/js/login_neu.js
	data := url.Values{}
	data.Set("rsa", account.GetRSA(ltID))
	data.Set("ul", strconv.Itoa(account.GetUserNameLength()))
	data.Set("pl", strconv.Itoa(account.GetPasswordLength()))
	data.Set("lt", ltID)
	data.Set("execution", executionID)
	data.Set("_eventId", "submit")
	req, _ = http.NewRequest("POST", config.DefaultGatewayURL, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", fmt.Sprintf("%d", len(data.Encode())))
	req.Header.Add("Referer", config.DefaultGatewayURL)
	resp, err = h.client.Do(req)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
		return err
	}
	span := doc.Find("span#errormsghide")
	text := span.Text()
	if len(text) != 0 {
		file, line := utils.GetErrorLocation()
		h.logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, text))
		return errors.New(text)
	}
	return nil
}

func getLtAndExecution(resp *http.Response) (string, string, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}
	lt := doc.Find("#lt").AttrOr("value", "")
	if len(lt) == 0 {
		return "", "", errors.New("lt not found")
	}
	execution := doc.Find("input[name='execution']").AttrOr("value", "")
	if len(execution) == 0 {
		return "", "", errors.New("execution not found")
	}
	return lt, execution, nil
}
