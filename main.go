package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type apiResponse struct {
	IP string `json:"ip"`
}

const apiEndpoint = "https://api.ipify.org/?format=json"

var currentIP = ""

var config jsonConfig

func main() {
	readSTDIN := true
	stat, _ := os.Stdin.Stat()
	if stat.Size() == 0 {
		log.Print("[WARNING] No configuration provided for ipaddresser from STDIN")
		config = defaultConfig()
		readSTDIN = false
	}

	if readSTDIN {
		jsonString, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Printf("[WARNING] Unable to read configuration for ipaddresser from STDIN: %s", err.Error())
			config = defaultConfig()
		}

		err = json.Unmarshal(jsonString, &config)
		if err != nil {
			log.Printf("[ERROR] Unable to parse config for ipaddresser: %s", err.Error())
			config = defaultConfig()
		}
	}

	if config.Daemon && config.QueryInterval < 1 {
		log.Print("[WARNING] Query interval for checking IP address must be 1 minute or greater; ignoring specified interval")
		config.QueryInterval = 1
	}

	if config.Daemon {
		monitoredQuery()
	} else {
		singleQuery()
	}
}

func getIPAddress() apiResponse {
	log.Printf("[INFO] Querying ipify at %s", apiEndpoint)

	res, err := http.Get(apiEndpoint)
	if err != nil {
		log.Printf("[ERROR] Unable to query ipify: %s", err.Error())
		os.Exit(1)
	}

	log.Print("[INFO] Parsing ipify response")

	var r apiResponse
	jsonString, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(jsonString), &r)
	if err != nil {
		log.Printf("[ERROR] Invalid response received from ipify: %s", jsonString)
		os.Exit(2)
	}

	return r
}
