package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

// Set configures slog's default logger (used by slog.Info, slog.Error, ...).
func Set(cfg Config) {
	slog.SetDefault(configure(cfg))
}

func configure(cfg Config) *slog.Logger {
	out := getOutput(cfg.Output)

	opts := &slog.HandlerOptions{
		AddSource:   cfg.ReportCaller,
		Level:       parseLevel(cfg.Level),
		ReplaceAttr: replaceAttr,
	}

	var handler slog.Handler
	switch cfg.Format {
	case FormatJSON:
		handler = slog.NewJSONHandler(out, opts)
	case FormatText:
		handler = slog.NewTextHandler(out, opts)
	default:
		handler = slog.NewJSONHandler(out, opts)
	}

	if cfg.Service != "" {
		handler = &serviceHandler{h: handler, service: cfg.Service}
	}

	return slog.New(handler)
}

func getOutput(out io.Writer) io.Writer {
	if out == nil {
		return os.Stdout
	}
	return out
}

func parseLevel(level string) slog.Leveler {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug
	case "info", "":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		if t, ok := a.Value.Any().(time.Time); ok {
			a.Value = slog.StringValue(t.Format("2006-01-02T15:04:05.999Z07:00"))
		}
		a.Key = "timestamp"
		return a
	case slog.LevelKey:
		if lv, ok := a.Value.Any().(slog.Level); ok {
			a.Value = slog.StringValue(strings.ToLower(lv.String()))
		}
		a.Key = "severity"
		return a
	case slog.MessageKey:
		a.Key = "message"
		return a
	case slog.SourceKey:
		src, ok := a.Value.Any().(*slog.Source)
		if !ok || src == nil {
			a.Key = "caller"
			return a
		}

		caller := fmt.Sprintf("%s:%d", src.File, src.Line)
		if src.Function != "" {
			parts := strings.Split(src.Function, ".")
			fn := parts[len(parts)-1]
			caller = fmt.Sprintf("%s() %s", fn, caller)
		}

		return slog.String("caller", caller)
	default:
		return a
	}
}

type serviceHandler struct {
	h       slog.Handler
	service string
}

func (sh *serviceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return sh.h.Enabled(ctx, level)
}

func (sh *serviceHandler) Handle(ctx context.Context, r slog.Record) error {
	if sh.service == "" {
		return sh.h.Handle(ctx, r)
	}

	hasService := false
	r.Attrs(func(a slog.Attr) bool {
		if a.Key == "service" {
			hasService = true
			return false
		}
		return true
	})

	if !hasService {
		r.AddAttrs(slog.String("service", sh.service))
	}
	return sh.h.Handle(ctx, r)
}

func (sh *serviceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &serviceHandler{h: sh.h.WithAttrs(attrs), service: sh.service}
}

func (sh *serviceHandler) WithGroup(name string) slog.Handler {
	return &serviceHandler{h: sh.h.WithGroup(name), service: sh.service}
}
