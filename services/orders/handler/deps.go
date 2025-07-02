package handler

import "log/slog"

type HandlerDeps struct {
	Log           *slog.Logger
	OrdersService OrdersService
}
