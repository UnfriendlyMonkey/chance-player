package shuffle_test

import (
	"testing"

	"github.com/unmo/chance-player/internal/domain/shuffle"
)

// fixedSource is a deterministic RandomSource for tests.
type fixedSource struct{ val int }

func (f fixedSource) Intn(_ int) int { return f.val }

func TestCoinPickHeads(t *testing.T) {
	c := shuffle.Coin{}
	got := c.Pick(fixedSource{0})
	if got != shuffle.Heads {
		t.Errorf("expected Heads, got %q", got)
	}
}

func TestCoinPickTails(t *testing.T) {
	c := shuffle.Coin{}
	got := c.Pick(fixedSource{1})
	if got != shuffle.Tails {
		t.Errorf("expected Tails, got %q", got)
	}
}

func TestDicePickRange(t *testing.T) {
	d := shuffle.NewDice(6)
	for i := 0; i < 6; i++ {
		got := d.Pick(fixedSource{i})
		want := i + 1
		if got != want {
			t.Errorf("Intn(%d) → want %d, got %d", i, want, got)
		}
	}
}

func TestDiceSides(t *testing.T) {
	d := shuffle.NewDice(20)
	if d.Sides() != 20 {
		t.Errorf("expected 20 sides, got %d", d.Sides())
	}
}

func TestCardDeckSize(t *testing.T) {
	deck := shuffle.NewStandard36Deck()
	// Verify via picks: 36 distinct cards
	seen := map[string]bool{}
	for i := 0; i < 36; i++ {
		card := deck.Pick(fixedSource{i})
		key := string(card.Suit) + string(card.Value)
		seen[key] = true
	}
	// All 36 picks via sequential indices should cover the deck
	if len(seen) != 36 {
		t.Errorf("expected 36 distinct cards, got %d", len(seen))
	}
}
