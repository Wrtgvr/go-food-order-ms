package app

import (
	"log/slog"
	"os"
)

func initLogger() *slog.Logger {
	// TODO: get env from config (when config will be added) and prepare config depends on env (local, dev, prod)
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
