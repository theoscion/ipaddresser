package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Ipify is the struct for interacting with the Ipify service
type Ipify struct {
	endpoint string
}

// GetIPAddress submits a GET request to Ipify to get the current IP address
func (ipify *Ipify) GetIPAddress(log *logrus.Logger) (string, error) {
	// Submit a GET request to Ipify to get the IP address
	log.Tracef("Querying Ipify at %s", ipify.endpoint)
	resp, err := http.Get(ipify.endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log the status code from the IP service response
	log.Debugf("Ipify responded with a %d status code", resp.StatusCode)

	// Parse the API response and return an error if it fails
	var apiResponse struct {
		IP string `json:"ip"`
	}
	jsonString, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(jsonString), &apiResponse)
	if err != nil {
		return "", err
	}

	// Return the IP, without error
	return apiResponse.IP, nil
}

// NewIpify returns an instance of the Ipify struct
func NewIpify() *Ipify {
	return &Ipify{
		endpoint: "https://api.ipify.org/?format=json",
	}
}
