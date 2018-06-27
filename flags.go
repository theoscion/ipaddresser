package main

import "flag"

var verbose = flag.Bool("v", false, "Specifies if logs should be verbose")

var daemon = flag.Bool("d", false, "Specifies to run as a daemon to watch for IP address changes")

var interval = flag.Int("i", 1, "Specifies how often the daemon should check for IP address changes; defaults to 1 minute")

var logOutput = flag.String("log", "", "Specifies an optional file to output all logs to instead of STDOUT")

var currentOutput = flag.String("out", "", "Specifies an optional file to output the current IP address to, as plain text")

var snsConfig = flag.String("sns", "", "Specifies an optional JSON configuration to use AWS SNS for IP address change notifications")

var emailConfig = flag.String("email", "", "Specifies an optional JSON configuration to use an email for IP address change notifications")
