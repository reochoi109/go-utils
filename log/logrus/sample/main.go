package main

import (
	ulog "utils/log/logrus/logger"

	log "github.com/sirupsen/logrus"
)

func main() {
	ulog.Set(ulog.PresetDev("mockup-service"))

	log.Info("INFO")
	log.Debug("DEBU")
	log.Warn("WARN")
}
