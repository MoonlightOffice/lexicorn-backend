package logger

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

var (
	Info  = logger.Info
	Warn  = logger.Warn
	Error = logger.Error
)
