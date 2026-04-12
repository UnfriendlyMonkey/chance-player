package usecase

import (
	"context"
	"fmt"

	"github.com/unmo/chance-player/internal/domain/chance"
	"github.com/unmo/chance-player/internal/domain/shuffle"
)

// chanceService is the concrete implementation of ChanceService.
// It holds all domain primitives and the RandomSource.
type chanceService struct {
	src  shuffle.RandomSource
	coin shuffle.Coin
	deck shuffle.CardDeck
}

// NewChanceService constructs a ChanceService backed by the given RandomSource.
func NewChanceService(src shuffle.RandomSource) ChanceService {
	return &chanceService{
		src:  src,
		coin: shuffle.Coin{},
		deck: shuffle.NewStandard36Deck(),
	}
}

func (s *chanceService) FlipCoin(_ context.Context) (chance.CoinFlipResult, error) {
	return chance.CoinFlipResult{Side: s.coin.Pick(s.src)}, nil
}

func (s *chanceService) FlipCoins(_ context.Context, count int) (chance.MultiCoinFlipResult, error) {
	if count < 1 {
		return chance.MultiCoinFlipResult{}, fmt.Errorf("count must be at least 1")
	}
	sides := make([]shuffle.Side, count)
	for i := range sides {
		sides[i] = s.coin.Pick(s.src)
	}
	return chance.MultiCoinFlipResult{Sides: sides, Count: count}, nil
}

func (s *chanceService) ConsultIChing(_ context.Context) (chance.IChingConsultResult, error) {
	var result chance.IChingConsultResult
	for i := range result.Lines {
		coins := make([]shuffle.Side, 3)
		sum := 0
		for j := range coins {
			coins[j] = s.coin.Pick(s.src)
			if coins[j] == shuffle.Heads {
				sum += 3
			} else {
				sum += 2
			}
		}
		result.Lines[i] = chance.IChingLineResult{
			Coins:    coins,
			Sum:      sum,
			LineType: iChingLineType(sum),
		}
	}
	return result, nil
}

func iChingLineType(sum int) chance.IChingLineType {
	switch sum {
	case 6:
		return chance.OldYin
	case 7:
		return chance.YoungYang
	case 8:
		return chance.YoungYin
	case 9:
		return chance.OldYang
	default:
		return chance.YoungYin
	}
}

func (s *chanceService) RollDice(_ context.Context, sides int) (chance.DiceRollResult, error) {
	if sides < 2 {
		return chance.DiceRollResult{}, fmt.Errorf("sides must be at least 2")
	}
	d := shuffle.NewDice(sides)
	return chance.DiceRollResult{Value: d.Pick(s.src), Sides: sides}, nil
}

func (s *chanceService) RollMultipleDice(_ context.Context, count, sides int) (chance.MultiDiceRollResult, error) {
	if count < 1 {
		return chance.MultiDiceRollResult{}, fmt.Errorf("count must be at least 1")
	}
	if sides < 2 {
		return chance.MultiDiceRollResult{}, fmt.Errorf("sides must be at least 2")
	}
	d := shuffle.NewDice(sides)
	rolls := make([]chance.DiceRollResult, count)
	sum := 0
	for i := range rolls {
		v := d.Pick(s.src)
		rolls[i] = chance.DiceRollResult{Value: v, Sides: sides}
		sum += v
	}
	return chance.MultiDiceRollResult{Rolls: rolls, Sum: sum, Sides: sides}, nil
}

func (s *chanceService) DrawCard(_ context.Context) (chance.CardDrawResult, error) {
	card := s.deck.Pick(s.src)
	return chance.CardDrawResult{Suit: card.Suit, Value: card.Value}, nil
}

func (s *chanceService) RandomNumber(_ context.Context, min, max int) (chance.RandomNumberResult, error) {
	if min >= max {
		return chance.RandomNumberResult{}, fmt.Errorf("min must be less than max")
	}
	n := s.src.Intn(max-min+1) + min
	return chance.RandomNumberResult{Number: n, Min: min, Max: max}, nil
}
