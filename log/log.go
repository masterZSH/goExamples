package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	f := &log.JSONFormatter{}
	f.TimestampFormat = "YYYY"

	log.SetFormatter(f)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	out, _ := os.OpenFile("1.log", os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(out)

	// Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

}
