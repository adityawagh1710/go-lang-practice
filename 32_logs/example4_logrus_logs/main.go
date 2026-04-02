package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"user": "aditya",
		"id":   101,
	}).Info("User logged in")

	log.Error("Something failed")
}
