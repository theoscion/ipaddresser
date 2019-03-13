package main

import "github.com/sirupsen/logrus"

// IPService represents the interface that needs to be satisfied to be used as an IP service
type IPService interface {
	GetIPAddress(*logrus.Logger) (string, error)
}
