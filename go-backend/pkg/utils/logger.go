package utils

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

func GetLogger() *slog.Logger {
	llev := slog.LevelInfo
	if os.Getenv("DEBUG") == "1" {
		llev = slog.LevelDebug
	}
	lopts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: llev,
		},
	}
	lhandler := NewPrettyHandler(os.Stdout, lopts)
	return slog.New(lhandler)
}

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

// Colorized slog logger
type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	// b, err := json.MarshalIndent(fields, "", "  ")
	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
