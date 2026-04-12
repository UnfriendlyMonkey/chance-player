// Package usecase defines the application use-case layer.
// All transport layers (HTTP, Telegram, future WebSocket) depend only on
// the ChanceService interface defined here — never on concrete implementations.
package usecase

import (
	"context"

	"github.com/unmo/chance-player/internal/domain/chance"
)

// ChanceService is the single input port for all randomization operations.
// It is the seam between transport and business logic.
type ChanceService interface {
	FlipCoin(ctx context.Context) (chance.CoinFlipResult, error)
	FlipCoins(ctx context.Context, count int) (chance.MultiCoinFlipResult, error)
	ConsultIChing(ctx context.Context) (chance.IChingConsultResult, error)

	RollDice(ctx context.Context, sides int) (chance.DiceRollResult, error)
	RollMultipleDice(ctx context.Context, count, sides int) (chance.MultiDiceRollResult, error)

	DrawCard(ctx context.Context) (chance.CardDrawResult, error)

	RandomNumber(ctx context.Context, min, max int) (chance.RandomNumberResult, error)
}
