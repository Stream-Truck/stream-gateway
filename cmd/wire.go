//go:build wireinject
// +build wireinject

package main

import (
	"application/config"
	http_v1 "application/internal/v1/http"
	"application/internal/v1/http/handler"
	"context"
	"github.com/google/wire"
	"log/slog"
	"net/http"
)

func wireApp(ctx context.Context, cfg config.Config, logger *slog.Logger) (http.Handler, error) {
	panic(wire.Build(
		http_v1.ServerProviderSet,
		handler.HandlerProviderSet,
	))
}
