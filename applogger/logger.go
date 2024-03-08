package applogger

import (
	"github.com/PaluMacil/dan2/config"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
)

func NewLogger(appEnv config.AppEnv) *slog.Logger {
	var handler slog.Handler
	if appEnv.LoggerUseStructured {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		// TODO: Color support on Windows can be added by using e.g. the go-colorable package
		handler = tint.NewHandler(os.Stderr, nil)
	}
	return slog.New(handler)
}
