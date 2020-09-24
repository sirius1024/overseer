package sniff

import (
	"os"
	"strings"

	"github.com/sirius1024/overseer/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// SelfReport if alive
func SelfReport() {
	conf := config.GetConfig()
	// Write to NSQ
	log.WithFields(log.Fields{
		"from": strings.Join([]string{conf.Cloud, conf.NetworkZone, conf.PrivateIP}, "-"),
		"to":   "-",
		"type": "self",
	}).Info("I'm alive!")
}
