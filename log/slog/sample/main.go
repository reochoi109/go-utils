package main

import (
	"log/slog"

	ulog "utils/log/slog/logger"
)

func main() {
	ulog.Set(ulog.PresetDev("mockup-service"))

	slog.Info("INFO", "request_id", "req-001")
	slog.Debug("DEBU", "request_id", "req-001")
	slog.Warn("WARN", "request_id", "req-001")
}
