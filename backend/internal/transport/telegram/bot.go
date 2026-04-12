package telegram

import (
	"context"
	"log/slog"
	"time"

	tele "gopkg.in/telebot.v3"

	"github.com/unmo/chance-player/internal/usecase"
	"github.com/unmo/chance-player/pkg/config"
)

// Bot wraps a telebot.Bot with the ChanceService.
type Bot struct {
	bot *tele.Bot
	svc usecase.ChanceService
}

// New creates and configures the Telegram bot.
func New(svc usecase.ChanceService, cfg *config.Config) (*Bot, error) {
	tb, err := tele.NewBot(tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	b := &Bot{bot: tb, svc: svc}
	b.registerHandlers()
	return b, nil
}

// Start begins polling for Telegram updates. Blocks until stopped.
func (b *Bot) Start() {
	slog.Info("telegram bot starting")
	b.bot.Start()
}

func (b *Bot) registerHandlers() {
	b.bot.Handle("/start", b.handleHelp)
	b.bot.Handle("/help", b.handleHelp)
	b.bot.Handle("/coin", b.handleCoin)
	b.bot.Handle("/3c", b.handle3Coins)
	b.bot.Handle("/iching", b.handleIChing)
	b.bot.Handle("/dice", b.handleDice)
	b.bot.Handle("/2d", b.handle2Dice)
	b.bot.Handle("/3d", b.handle3Dice)
	b.bot.Handle("/card", b.handleCard)
	b.bot.Handle("/number", b.handleNumber)
}

func (b *Bot) handleHelp(c tele.Context) error {
	return c.Send(`🎲 Chance Player Bot

Commands:
/coin — flip a coin
/3c — flip 3 coins
/iching — I Ching consultation (6 lines)
/dice — roll a d6
/2d — roll 2d6
/3d — roll 3d6
/card — draw a card
/number — random number 1–100`)
}

func (b *Bot) handleCoin(c tele.Context) error {
	r, err := b.svc.FlipCoin(context.Background())
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatCoinFlip(r))
}

func (b *Bot) handle3Coins(c tele.Context) error {
	r, err := b.svc.FlipCoins(context.Background(), 3)
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatMultiCoin(r))
}

func (b *Bot) handleIChing(c tele.Context) error {
	r, err := b.svc.ConsultIChing(context.Background())
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatIChing(r))
}

func (b *Bot) handleDice(c tele.Context) error {
	r, err := b.svc.RollDice(context.Background(), 6)
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatDiceRoll(r))
}

func (b *Bot) handle2Dice(c tele.Context) error {
	r, err := b.svc.RollMultipleDice(context.Background(), 2, 6)
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatMultiDice(r))
}

func (b *Bot) handle3Dice(c tele.Context) error {
	r, err := b.svc.RollMultipleDice(context.Background(), 3, 6)
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatMultiDice(r))
}

func (b *Bot) handleCard(c tele.Context) error {
	r, err := b.svc.DrawCard(context.Background())
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatCard(r))
}

func (b *Bot) handleNumber(c tele.Context) error {
	r, err := b.svc.RandomNumber(context.Background(), 1, 100)
	if err != nil {
		return c.Send("error: " + err.Error())
	}
	return c.Send(formatNumber(r))
}
