package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// main is the initializing function for the app
func main() {
	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	cli := NewCLI(args)
	cli.Run(logrus.New(), NewIpify())
}
