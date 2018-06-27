package main

import (
	"fmt"
	"log"
	"time"
)

func Get() string {
	r := getIPAddress()
	return r.IP
}

func singleQuery() {
	r := getIPAddress()

	log.Printf("[RESULT] Current IP Address\t%s", r.IP)

	fmt.Println(r.IP)
}

func monitoredQuery() {
	for {
		r := getIPAddress()

		if currentIP == "" {
			log.Printf("[RESULT] Current IP Address:\t%s", r.IP)
			fmt.Println(r.IP)
			runHook(r.IP, true)
		} else if r.IP != currentIP {
			log.Printf("[RESULT] New IP Address:t%s", r.IP)
			fmt.Println(r.IP)
			runHook(r.IP, false)
		} else {
			log.Printf("[RESULT] Same IP Address:\t%s", r.IP)
		}

		currentIP = r.IP

		if config.QueryInterval == 1 {
			log.Print("[INFO] Sleeping for 1 minute")
		} else {
			log.Printf("[INFO] Sleeping for %d minutes", config.QueryInterval)
		}

		time.Sleep(time.Duration(config.QueryInterval) * time.Minute)

		log.Println("[INFO] Waking up from sleep")
	}
}
