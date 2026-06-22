package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/DoraTiger/NEU_IPGW/config"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/sirupsen/logrus"
)

type ElectricityHandler struct {
	client *http.Client
	logger *logrus.Logger
	sso    *SSOHandler
}

type RoomInfo struct {
	LastPayFlag   string `json:"lastpayflag"`
	ElecRoomNo    string `json:"elecroomno"`
	DormitoryName string `json:"-"`
}

type ElectricityInfo struct {
	ElecRoomInfo string  `json:"elecroominfo"`
	ElecRemain   string  `json:"elecRemain"`
	FacFlag      string  `json:"facFlag"`
	MAddr        string  `json:"mAddr"`
	LastReadTime string  `json:"lastReadTime"`
	ReturnCode   string  `json:"returncode"`
	ReturnMsg    string  `json:"returnmsg"`
	RemainDegree float64 `json:"-"`
	RemainYuan   float64 `json:"-"`
}

func NewElectricityHandler() *ElectricityHandler {
	client := utils.NewClientWithJar()
	return &ElectricityHandler{
		client: client,
		sso:    NewSSOHandler(client),
	}
}

func (h *ElectricityHandler) SetLogger(logger *logrus.Logger) {
	h.logger = logger
	h.sso.SetLogger(logger)
}

func (h *ElectricityHandler) SetClient(client *http.Client) {
	h.client = client
	h.sso.client = client
}

func (h *ElectricityHandler) Login() error {
	if err := h.sso.Login(config.DefaultSSOBaseURL); err != nil {
		return fmt.Errorf("CAS login: %w", err)
	}

	resp, err := h.client.Get(config.DefaultPayCASURL)
	if err != nil {
		return fmt.Errorf("pay CAS login: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("pay CAS login HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (h *ElectricityHandler) QueryDefaultRoomInfo(factoryCode string) (*RoomInfo, error) {
	form := url.Values{}
	form.Set("factorycode", factoryCode)

	req, err := http.NewRequest("POST", config.DefaultElectricityQueryURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", config.DefaultPayReferer)

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var roomInfo RoomInfo
	if err := utils.ParseJSON(body, &roomInfo); err != nil {
		return nil, fmt.Errorf("parse response: %w (raw: %s)", err, string(body))
	}

	if roomInfo.ElecRoomNo == "" {
		return nil, fmt.Errorf("no room number found for this account, please specify with -r flag")
	}

	roomInfo.DormitoryName = config.TranslateDormitory(roomInfo.ElecRoomNo)

	return &roomInfo, nil
}

func (h *ElectricityHandler) Logout() error {
	return h.sso.Logout(config.DefaultSSOLogoutURL)
}

func (h *ElectricityHandler) QueryRoomInfo(elecRoomNo, factoryCode string) (*ElectricityInfo, error) {
	form := url.Values{}
	form.Set("elecroomno", elecRoomNo)
	form.Set("factorycode", factoryCode)

	req, err := http.NewRequest("POST", config.DefaultElectricityRoomInfoURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", config.DefaultPayReferer)

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var info ElectricityInfo
	if err := utils.ParseJSON(body, &info); err != nil {
		return nil, fmt.Errorf("parse response: %w (raw: %s)", err, string(body))
	}

	if info.ReturnCode != "SUCCESS" {
		return nil, fmt.Errorf("API error: %s - %s", info.ReturnCode, info.ReturnMsg)
	}

	info.parseElecRemain()

	return &info, nil
}

func (info *ElectricityInfo) parseElecRemain() {
	var value float64
	var isYuan bool

	fmt.Sscanf(info.ElecRemain, "%f", &value)

	if strings.Contains(info.ElecRemain, "元") {
		isYuan = true
	}

	if isYuan {
		info.RemainYuan = value
		info.RemainDegree = value / config.DefaultElectricityPrice
	} else {
		info.RemainDegree = value
		info.RemainYuan = value * config.DefaultElectricityPrice
	}
}
