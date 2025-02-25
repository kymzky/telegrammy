package utils

import (
	"log/slog"
	"os"
)

func InitializeLogger() {
	handler := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(handler)
}
