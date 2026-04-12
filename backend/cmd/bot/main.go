package main

import (
	"log/slog"
	"os"

	"github.com/unmo/chance-player/internal/infra/random"
	"github.com/unmo/chance-player/internal/transport/telegram"
	"github.com/unmo/chance-player/internal/usecase"
	"github.com/unmo/chance-player/pkg/config"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if cfg.BotToken == "" {
		slog.Error("BOT_TOKEN environment variable is required")
		os.Exit(1)
	}

	src := random.CryptoSource{}
	svc := usecase.NewChanceService(src)

	bot, err := telegram.New(svc, cfg)
	if err != nil {
		slog.Error("failed to create bot", "err", err)
		os.Exit(1)
	}

	bot.Start()
}
