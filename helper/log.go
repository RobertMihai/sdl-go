package helper

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func InitLogger() {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func GetLogger() *slog.Logger {
	return log
}
