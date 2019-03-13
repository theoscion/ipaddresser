package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// CLI is a struct that is used for running the application and has various configuration settings
type CLI struct {
	Verbose    bool
	Daemon     bool
	Hooks      map[string]int
	AlwaysHook bool
}

// Run will run the daemon
func (cli *CLI) Run(log *logrus.Logger, ipService IPService) {
	// Set the log level based on if verbose logging is enabled or not
	if cli.Verbose {
		log.SetLevel(logrus.TraceLevel)
		log.Trace("Showing verbose output")
	} else {
		log.SetLevel(logrus.FatalLevel)
	}

	// Log how many webhooks are registered
	hookLen := len(cli.Hooks)
	hookNoun := "webhooks"
	if hookLen == 1 {
		hookNoun = "webhook"
	}
	log.Debugf("Registered %d %s", hookLen, hookNoun)

	// Run the application in daemon mode, if specified; otherwise, do a single request
	if cli.Daemon {
		// Log the we are running in daemon mode
		log.Trace("Running in daemon mode")

		// Prepare the current IP and first check variables
		currentIP := ""
		firstCheck := true

		// Run in an infinite loop, until the user breaks
		for {
			// Retrieve the IP address from the IP service
			log.Trace("Querying IP service to retrieve IP address")
			ip, err := ipService.GetIPAddress(log)
			if err != nil {
				log.Error(err)
				fmt.Println("Failed getting IP address from IP service")
				return
			}

			// Handle any changes to the IP address
			if currentIP == "" {
				// Log the current IP
				log.WithField("ip", ip).Info("Determined current IP address")
				fmt.Printf("Current IP address set to %s\n", ip)
			} else if ip != currentIP {
				// Log the current IP
				log.WithField("ip", ip).Info("Determined new IP address")
				fmt.Printf("New IP address set to %s\n", ip)
			} else {
				log.WithField("ip", ip).Info("No change to IP address")
				fmt.Printf("IP address is unchanged\n")
			}

			// Run any webhooks
			log.Trace("Running defined webhooks")
			cli.RunAllWebhooks(log, currentIP, ip, firstCheck)

			// Set the current IP
			currentIP = ip

			// Set that this is no longer the first IP check
			firstCheck = false

			// Sleep the daemon for 1 minute
			log.Trace("Going to sleep for 1 minute")
			time.Sleep(time.Minute)
			log.Trace("Waking up to check for IP address changes")
		}
	} else {
		// Retrieve the IP address from the IP service
		log.Trace("Querying IP service to retrieve IP address")
		ip, err := ipService.GetIPAddress(log)
		if err != nil {
			log.Error(err)
			fmt.Println("Failed getting IP address from IP service")
			return
		}

		// Run any webhooks
		log.Trace("Running defined webhooks")
		cli.RunAllWebhooks(log, "", ip, true)
		log.Trace("All webhooks completed")

		// Output the IP address from the IP service
		log.WithField("ip", ip).Debug("IP address retrieved successfully")
		fmt.Println(ip)
	}
}

// RunAllWebhooks handles running all CLI webhooks, if specified
func (cli *CLI) RunAllWebhooks(log *logrus.Logger, oldIP string, newIP string, firstCheck bool) {
	// If there are no hook, return without executing
	if len(cli.Hooks) == 0 {
		log.Debug("No webhooks to run")
		return
	}

	// If the IP hasn't changed and the hook shouldn't run, return without executing
	if oldIP == newIP && !cli.AlwaysHook {
		log.Debug("All webhooks skipped because IP address is unchanged")
		return
	}

	// Prepare the hook data
	var hookData struct {
		Source     string `json:"source"`
		FirstCheck bool   `json:"firstCheck"`
		OldIP      string `json:"oldIP"`
		NewIP      string `json:"newIP"`
	}
	hookData.Source = "ipaddresser"
	hookData.FirstCheck = firstCheck
	hookData.OldIP = oldIP
	hookData.NewIP = newIP
	jsonData, _ := json.Marshal(hookData)

	// Loop through all webhooks and submit an HTTP request for each webhook
	var hookSync sync.WaitGroup
	for hook := range cli.Hooks {
		hookSync.Add(1)
		go cli.RunWebhook(log, &hookSync, hook, jsonData)
	}
	hookSync.Wait()
}

// RunWebhook handles running a CLI webhook concurrently
func (cli *CLI) RunWebhook(log *logrus.Logger, hookSync *sync.WaitGroup, hook string, jsonData []byte) {
	// Build the webhook request
	log.WithField("url", hook).Trace("Submitting webhook request")
	req, err := http.NewRequest("POST", hook, bytes.NewBuffer(jsonData))
	req.Header.Set("User-Agent", "ForceCore/ipaddresser")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.WithField("url", hook).Errorf("Hook request creation failed: %s", err.Error())
		hookSync.Done()
		return
	}

	// Submit the webhook request
	log.WithField("url", hook).Tracef("Calling webhook using POST")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("url", hook).Errorf("Webhook request failed: %s", err)
		hookSync.Done()
		return
	}
	defer resp.Body.Close()

	// Log the status code from the IP service response
	log.WithField("url", hook).Debugf("Webhook responded with a %d status code", resp.StatusCode)

	// Mark this call as done in the hook sync
	hookSync.Done()
}

// NewCLI returns an instance of the CLI struct, with flags parsed
func NewCLI(args []string) *CLI {
	// Create a CLI instance
	cli := CLI{}

	// Prepare the list of hooks
	cli.Hooks = make(map[string]int)

	// Loop through all arguments and process any flags or hooks
	for _, arg := range args {
		if arg == "--verbose" || arg == "-v" {
			cli.Verbose = true
		} else if arg == "--daemon" || arg == "-d" {
			cli.Daemon = true
		} else if arg == "--always-hook" || arg == "-a" {
			cli.AlwaysHook = true
		} else if arg[0:1] != "-" {
			cli.Hooks[arg] = 0
		}
	}

	// Return the CLI instance as a pointer
	return &cli
}
