package usecase_test

import (
	"context"
	"testing"

	"github.com/unmo/chance-player/internal/domain/shuffle"
	"github.com/unmo/chance-player/internal/usecase"
)

// seqSource cycles through values: 0,1,2,3,...
type seqSource struct{ calls int }

func (s *seqSource) Intn(n int) int {
	v := s.calls % n
	s.calls++
	return v
}

// fixedSource always returns the same value.
type fixedSource struct{ val int }

func (f fixedSource) Intn(_ int) int { return f.val }

func TestFlipCoin(t *testing.T) {
	svc := usecase.NewChanceService(fixedSource{0})
	r, err := svc.FlipCoin(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if r.Side != shuffle.Heads {
		t.Errorf("expected heads, got %q", r.Side)
	}
}

func TestFlipCoinsCount(t *testing.T) {
	svc := usecase.NewChanceService(&seqSource{})
	r, err := svc.FlipCoins(context.Background(), 3)
	if err != nil {
		t.Fatal(err)
	}
	if r.Count != 3 || len(r.Sides) != 3 {
		t.Errorf("expected 3 coins, got count=%d len=%d", r.Count, len(r.Sides))
	}
}

func TestFlipCoinsInvalidCount(t *testing.T) {
	svc := usecase.NewChanceService(fixedSource{0})
	_, err := svc.FlipCoins(context.Background(), 0)
	if err == nil {
		t.Error("expected error for count=0")
	}
}

func TestRollDice(t *testing.T) {
	svc := usecase.NewChanceService(fixedSource{2}) // Intn(6)=2 → value=3
	r, err := svc.RollDice(context.Background(), 6)
	if err != nil {
		t.Fatal(err)
	}
	if r.Value != 3 || r.Sides != 6 {
		t.Errorf("expected value=3 sides=6, got value=%d sides=%d", r.Value, r.Sides)
	}
}

func TestRollMultipleDice(t *testing.T) {
	svc := usecase.NewChanceService(&seqSource{}) // returns 0,1,2 → values 1,2,3 sum=6
	r, err := svc.RollMultipleDice(context.Background(), 3, 6)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Rolls) != 3 {
		t.Errorf("expected 3 rolls, got %d", len(r.Rolls))
	}
	if r.Sum != 6 {
		t.Errorf("expected sum=6, got %d", r.Sum)
	}
}

func TestDrawCard(t *testing.T) {
	svc := usecase.NewChanceService(fixedSource{0})
	r, err := svc.DrawCard(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if r.Suit == "" || r.Value == "" {
		t.Error("expected non-empty suit and value")
	}
}

func TestRandomNumber(t *testing.T) {
	svc := usecase.NewChanceService(fixedSource{0}) // Intn(91)=0 → 1+0=1
	r, err := svc.RandomNumber(context.Background(), 1, 91)
	if err != nil {
		t.Fatal(err)
	}
	if r.Number < 1 || r.Number > 91 {
		t.Errorf("number %d out of range [1,91]", r.Number)
	}
}

func TestRandomNumberInvalidRange(t *testing.T) {
	svc := usecase.NewChanceService(fixedSource{0})
	_, err := svc.RandomNumber(context.Background(), 10, 5)
	if err == nil {
		t.Error("expected error for min >= max")
	}
}

func TestConsultIChing(t *testing.T) {
	svc := usecase.NewChanceService(&seqSource{})
	r, err := svc.ConsultIChing(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for i, line := range r.Lines {
		if len(line.Coins) != 3 {
			t.Errorf("line %d: expected 3 coins, got %d", i, len(line.Coins))
		}
		if line.Sum < 6 || line.Sum > 9 {
			t.Errorf("line %d: sum %d out of I Ching range [6,9]", i, line.Sum)
		}
	}
}
