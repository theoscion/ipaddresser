package main

import (
	"fmt"
	"log"
	"time"
)

func monitoredQuery() {
	for {
		r := getIPAddress()

		changed := false

		if currentIP == "" {
			log.Printf("[RESULT] Current IP Address\t%s", r.IP)
			outputCurrentIP(r.IP)
		} else if r.IP != currentIP {
			log.Printf("[RESULT] New IP Address\t%s", r.IP)
			outputCurrentIP(r.IP)
			changed = true
		} else if *verbose {
			log.Printf("[RESULT] Same IP Address\t%s", r.IP)
		}

		currentIP = r.IP

		if changed {
			if *snsConfig != "" {
				sendSNSNotification()
			}

			if *emailConfig != "" {
				sendEmailNotification()
			}
		}

		if *verbose {
			log.Printf("[INFO] Sleeping for %d minute", *interval)
		}

		time.Sleep(time.Duration(*interval) * time.Minute)

		if *verbose {
			log.Println("[INFO] Waking up from sleep")
		}
	}
}

func singleQuery() {
	r := getIPAddress()
	outputCurrentIP(r.IP)

	if *verbose {
		log.Printf("[RESULT] Current IP Address\t%s", r.IP)
	} else {
		fmt.Println(r.IP)
	}
}
