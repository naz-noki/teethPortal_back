package logger

import (
	"errors"
	"log/slog"
	"os"
)

var (
	Log *slog.Logger
)

func Init(mode, pathToLogsFile string) (*os.File, error) {
	var l *slog.Logger

	file, errOpenFile := os.OpenFile(pathToLogsFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if errOpenFile != nil {
		file.Close()
		return nil, errOpenFile
	}

	switch mode {
	case "local":
		l = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		l = slog.New(
			slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		l = slog.New(
			slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		file.Close()
		return nil, errors.New("unknown logger mode")
	}

	Log = l
	return file, nil
}
