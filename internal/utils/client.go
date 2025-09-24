package utils

import (
	"crypto/tls"
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
