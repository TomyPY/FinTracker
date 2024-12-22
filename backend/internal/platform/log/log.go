package log

import (
	"log/slog"
	"os"
)

func SetLog() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
