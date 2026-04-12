package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/unmo/chance-player/internal/infra/random"
	transport "github.com/unmo/chance-player/internal/transport/http"
	"github.com/unmo/chance-player/internal/usecase"
	"github.com/unmo/chance-player/pkg/config"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	src := random.CryptoSource{}
	svc := usecase.NewChanceService(src)

	srv := transport.New(svc, cfg)

	addr := ":" + cfg.HTTPPort
	slog.Info("starting API server", "addr", addr, "env", cfg.Env)

	if err := http.ListenAndServe(addr, srv.Mux); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
