package main

import (
	"application/config"
	"application/internal/v1/biz"
	"application/internal/v1/datasource"
	http_v1 "application/internal/v1/http"
	"application/internal/v1/http/handler"
	"context"
	"github.com/google/wire"
	"log/slog"
	"net/http"
)

func wireApp(ctx context.Context, cfg config.Config, logger *slog.Logger) (http.Handler, error) {
	panic(wire.Build(
		datasource.DataProviderSet,
		biz.ProviderSet,
		http_v1.ServerProviderSet,
		handler.HandlerProviderSet,
	))
}
