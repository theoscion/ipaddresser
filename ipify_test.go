package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestIpify(test *testing.T) {
	ipify := NewIpify()
	assert.Equal(test, "https://api.ipify.org/?format=json", ipify.endpoint, "ipify.endpoint should match")
}

func TestIpifyRequest(test *testing.T) {
	ipify := NewIpify()

	var apiResponse struct {
		IP string `json:"ip"`
	}
	resp, _ := http.Get(ipify.endpoint)
	jsonString, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(jsonString), &apiResponse)

	ip, err := ipify.GetIPAddress(logrus.New())
	assert.Nil(test, err, "ipify.GetIPAddress() should not fail")
	assert.Equal(test, apiResponse.IP, ip, "ipify.GetIPAddress() should match")
}
