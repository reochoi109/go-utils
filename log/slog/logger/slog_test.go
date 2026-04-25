package logger

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Success Configuration", func(t *testing.T) {
		cfg := Config{
			Service:      "test-service",
			Format:       FormatJSON,
			Level:        "debug",
			ReportCaller: true,
		}

		log := configure(cfg)
		if log == nil {
			t.Fatal("expected non-nil logger")
		}
	})

	t.Run("Invalid Level Fallback", func(t *testing.T) {
		cfg := Config{Level: "invalid-level"}
		log := configure(cfg)
		if log == nil {
			t.Fatal("expected non-nil logger")
		}
	})
}

func TestGetOutput(t *testing.T) {
	var buf bytes.Buffer
	cfg := PresetDev("test")
	cfg.Output = io.MultiWriter(&buf, os.Stdout)

	log := configure(cfg)

	log.Info("hello info log")
	if buf.Len() == 0 {
		t.Error("not print log")
	}
}

func TestCustomCallerPrettyfier(t *testing.T) {
	var buf bytes.Buffer
	cfg := PresetDev("test")
	cfg.ReportCaller = true
	cfg.Output = &buf

	log := configure(cfg)
	log.Info("test message")

	output := buf.String()
	if !strings.Contains(output, "TestCustomCallerPrettyfier()") {
		t.Errorf("invalid print to Caller. Got: %s", output)
	}
}

func TestServiceHook(t *testing.T) {
	serviceName := "prod-service"
	var buf bytes.Buffer
	cfg := PresetProd(serviceName)
	cfg.Output = &buf

	log := configure(cfg)
	log.Info("hook test")

	output := buf.String()
	if !strings.Contains(output, serviceName) {
		t.Errorf("Service name missing in log. Expected to contain: %s, Got: %s", serviceName, output)
	}
}

func TestSet(t *testing.T) {
	var buf bytes.Buffer
	cfg := PresetProd("set-service")
	cfg.Format = FormatJSON
	cfg.Output = &buf

	Set(cfg)
	slog.Info("set test")

	if !strings.Contains(buf.String(), "set test") {
		t.Fatalf("expected output to include log message, got: %s", buf.String())
	}
}
