package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type hookData struct {
	IP         string `json:"ip"`
	FirstQuery bool   `json:"firstQuery"`
}

// runHook handles running the HTTP hook when an IP address is first loaded or changed
func runHook(ipAddress string, firstQuery bool) {
	// If the HTTP hook is enabled, validate and run the hook
	if config.Hook.Enabled {
		// Make sure there is a URL provided; if not disable the hook
		if config.Hook.URL == "" {
			config.Hook.Enabled = false
		}

		// If a hook method isn't provided, set it to GET
		if config.Hook.Method == "" {
			config.Hook.Method = "GET"
		}

		// If the hook is still enabled, run the hook
		if config.Hook.Enabled {
			// Prepare the data to send
			data := hookData{
				IP:         ipAddress,
				FirstQuery: firstQuery,
			}
			jsonData, _ := json.Marshal(data)

			// Create the HTTP request
			request, err := http.NewRequest(config.Hook.Method, config.Hook.URL, bytes.NewBuffer(jsonData))
			request.Header.Set("Content-Type", "application/json")
			if err != nil {
				log.Printf("[ERROR] Hook request creation failed: %s", err.Error())
				return
			}

			// Submit the request
			log.Printf("[INFO] Hook submitted to %s (%s)", config.Hook.URL, config.Hook.Method)
			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				log.Printf("[ERROR] Hook submission failed: %s", err.Error())
				return
			}
			defer response.Body.Close()

			// Log the result
			log.Printf("[INFO] Hook responded with: %s", response.Status)
		}
	}
}
