package main

import (
	"fmt"
	"log"
	"time"
)

// Get will return query ipify and return the resulting IP address as a string; exported for use in other packages
func Get() string {
	// Get the IP address from ipify
	r := getIPAddress()

	// Return the IP address
	return r.IP
}

// singleQuery query ipify and print to log and screen the result
func singleQuery() {
	// Get the IP address from ipify
	r := getIPAddress()

	// Log the IP address
	log.Printf("[RESULT] Current IP Address\t%s", r.IP)

	// Print the IP address
	fmt.Println(r.IP)
}

// monitorQuery will run as a daemon, querying ipify at the rate of the query interval (specified in config), and the results will be printed to log and screen
func monitoredQuery() {
	for {
		// Get the IP address from ipify
		r := getIPAddress()

		// Determine how to print the result, based on if the IP address has already been loaded, if it has changed, or if it is the same
		if currentIP == "" {
			log.Printf("[RESULT] Current IP Address:\t%s", r.IP)
			runHook(r.IP, true)
			fmt.Println(r.IP)
		} else if r.IP != currentIP {
			log.Printf("[RESULT] New IP Address:t%s", r.IP)
			runHook(r.IP, false)
			fmt.Println(r.IP)
		} else {
			log.Printf("[RESULT] Same IP Address:\t%s", r.IP)
			if config.Hook.SubmitForSameIP {
				runHook(r.IP, false)
			}
		}

		// Update the current IP address to the new IP address
		currentIP = r.IP

		// Log how long the app is going to sleep before next query
		if config.QueryInterval == 1 {
			log.Print("[INFO] Sleeping for 1 minute")
		} else {
			log.Printf("[INFO] Sleeping for %d minutes", config.QueryInterval)
		}

		// Put the app to sleep
		time.Sleep(time.Duration(config.QueryInterval) * time.Minute)

		// Log that the app has woken up
		log.Println("[INFO] Waking up from sleep")
	}
}
