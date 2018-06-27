package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// apiResponse is a struct that can parse the result from the ipify service
type apiResponse struct {
	IP string `json:"ip"` // Specifies the IP address
}

// apiEndpoint is a constant that specifies the fully-qualified endpoint for retrieving a public IP address
const apiEndpoint = "https://api.ipify.org/?format=json"

// currentIP contains the last known IP addressed returned from an ipify query; this is only used when running as a daemon
var currentIP = ""

// config contains the JSON configuration used when launched, either using defaultConfig() or a JSON string from STDIN
var config jsonConfig

// main is the initializing function for the app
func main() {
	// Specifies that STDIN should be read (used to determine if reading should be skipped due to no content)
	readSTDIN := true

	// Determine if there is content on STDIN; if not, log a warning, load the default config, and set readSTDIN to false
	stat, _ := os.Stdin.Stat()
	if stat.Size() == 0 {
		log.Print("[WARNING] No configuration provided for ipaddresser from STDIN")
		config = defaultConfig()
		readSTDIN = false
	}

	// If STDIN should be read, read and parse the STDIN as JSON. If it fails, log a warning and use the default config
	if readSTDIN {
		jsonString, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Printf("[WARNING] Unable to read configuration for ipaddresser from STDIN: %s", err.Error())
			config = defaultConfig()
		}

		err = json.Unmarshal(jsonString, &config)
		if err != nil {
			log.Printf("[WARNING] Unable to parse config for ipaddresser: %s", err.Error())
			config = defaultConfig()
		}
	}

	// If this app is running as a daemon and the query interval is less than one, log a warning and override the interval back to 1
	if config.Daemon && config.QueryInterval < 1 {
		log.Print("[WARNING] Query interval for checking IP address must be 1 minute or greater; ignoring specified interval")
		config.QueryInterval = 1
	}

	// Submit the queries based on if it is a single query or a monitored query (a.k.a., running as a daemon)
	if config.Daemon {
		monitoredQuery()
	} else {
		singleQuery()
	}
}

// getIPAddress handles submitting a query to ipify to get a public IP address
func getIPAddress() apiResponse {
	// Log info for querying ipify
	log.Printf("[INFO] Querying ipify at %s", apiEndpoint)

	// Submit a request to the ipify API endpoint; if it fails, log an error and exit with return code 1
	res, err := http.Get(apiEndpoint)
	if err != nil {
		log.Printf("[ERROR] Unable to query ipify: %s", err.Error())
		os.Exit(1)
	}

	// Log info for parsing the ipify response
	log.Print("[INFO] Parsing ipify response")

	// Parse the ipify JSON response into the apiResponse{} struct; if it fails, log an error and exist with return code 2
	var r apiResponse
	jsonString, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(jsonString), &r)
	if err != nil {
		log.Printf("[ERROR] Invalid response received from ipify: %s", jsonString)
		os.Exit(2)
	}

	// Return the API response
	return r
}
