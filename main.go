package main

import (
	"encoding/json"
	"flag"
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

func main() {
	flag.Parse()

	if *logOutput != "" {
		lf, err := os.Create(*logOutput)
		if err != nil {
			if *verbose {
				log.Printf("[WARNING] Unable to access log output file: %s", err.Error())
			}

			*logOutput = ""
		} else {
			log.SetOutput(lf)
		}
		defer lf.Close()
	}

	if *daemon && *interval < 1 {
		log.Print("[WARNING] Monitor interval must be at least 1 minute; ignoring specified interval")
		*interval = 1
	}

	if *daemon {
		monitoredQuery()
	} else {
		singleQuery()
	}
}

func getIPAddress() apiResponse {
	if *verbose {
		log.Printf("[INFO] Querying ipify at %s", apiEndpoint)
	}

	res, err := http.Get(apiEndpoint)
	if err != nil {
		log.Printf("[ERROR] Unable to query ipify: %s", err.Error())
		os.Exit(1)
	}

	if *verbose {
		log.Print("[INFO] Parsing ipify response")
	}

	var r apiResponse
	jsonString, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(jsonString), &r)
	if err != nil {
		log.Printf("[ERROR] Invalid response received from ipify: %s", jsonString)
		os.Exit(2)
	}

	return r
}

func outputCurrentIP(ipAddress string) {
	if *currentOutput != "" {
		if *verbose {
			log.Printf("[INFO] Writing current IP address to %s", *currentOutput)
		}

		writeFile := true

		f, err := os.Create(*currentOutput)
		if err != nil {
			if *verbose {
				log.Printf("[WARNING] Unable to access current output file: %s", err.Error())
			}

			*currentOutput = ""
			writeFile = false
		}
		defer f.Close()

		if writeFile {
			_, err = f.Write([]byte(ipAddress))
			if err != nil {
				if *verbose {
					log.Printf("[WARNING] Unable to write IP address to current output file: %s", err.Error())
				}

				*currentOutput = ""
			}
		}
	}
}
