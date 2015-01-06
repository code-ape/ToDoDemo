package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

const (
	DEBUG     = true
	PRINT_LOG = true
)

func ConfigLogging() {
	if PRINT_LOG {
		log.SetOutput(os.Stdout)
	}
	//log.SetFormatter(log.JSONFormatter)
	if DEBUG {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.Debug("Logging configured.")
}
