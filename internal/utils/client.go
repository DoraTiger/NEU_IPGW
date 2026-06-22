package utils

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
)

// create new client with cookiejar and set InsecureSkipVerify true
func NewClientWithJar() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return client
}

func ParseJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
