package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 添加命令行参数
	usernamePtr := flag.String("username", "", "the username for authentication")
	passwordPtr := flag.String("password", "", "the password for authentication")
	flag.Parse()

	// 检查命令行参数是否提供
	if *usernamePtr == "" || *passwordPtr == "" {
		fmt.Println("Error: both username and password are required")
		return
	}

	// 创建HTTP客户端
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 访问统一网关，获取cookie
	// https://pass.neu.edu.cn/tpass/login
	gatewayURL := "https://pass.neu.edu.cn/tpass/login"
	req, _ := http.NewRequest("GET", gatewayURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error visit gateway:", err)
		return
	}
	defer resp.Body.Close()

	ltID, executionID, err := getLtAndExecution(resp)
	if err != nil {
		fmt.Println("Error get lt:", err)
		return
	}

	// 参数定义参考 https://pass.neu.edu.cn/tpass/comm/neu/js/login_neu.js
	data := url.Values{}
	data.Set("rsa", fmt.Sprintf("%s%s%s", *usernamePtr, *passwordPtr, ltID))
	data.Set("ul", strconv.Itoa(len(*usernamePtr)))
	data.Set("pl", strconv.Itoa(len(*passwordPtr)))
	data.Set("lt", ltID)
	data.Set("execution", executionID)
	data.Set("_eventId", "submit")
	req, _ = http.NewRequest("POST", gatewayURL, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", fmt.Sprintf("%d", len(data.Encode())))
	req.Header.Add("Referer", gatewayURL)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error get cookie:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("can not read response:", err)
		return
	}
	span := doc.Find("span#errormsghide")
	text := span.Text()
	if len(text) != 0 {
		fmt.Println(text)
		return
	}

	// 登录联网网关
	// https://ipgw.neu.edu.cn/

	// 获取当前网络下对应网关url的query参数
	resp, err = client.Get("https://ipgw.neu.edu.cn/")
	if err != nil {
		fmt.Println("Error get ipgw address:", err)
		return
	}
	defer resp.Body.Close()
	// 统一认证拿到ticket
	resp, err = client.Get("https://pass.neu.edu.cn/tpass/login?service=http://ipgw.neu.edu.cn/srun_portal_sso?" + resp.Request.URL.RawQuery)
	fmt.Println(resp.Request.URL.RawQuery)
	if err != nil {
		fmt.Println("Error get ticket:", err)
		return
	}
	defer resp.Body.Close()
	// 使用ticket调用api登录
	req, _ = http.NewRequest("GET", "https://ipgw.neu.edu.cn/v1"+resp.Request.URL.RequestURI(), nil)
	fmt.Println(resp.Request.URL.RequestURI())
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error login ipgw:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(strconv.Quote(string(body)))

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