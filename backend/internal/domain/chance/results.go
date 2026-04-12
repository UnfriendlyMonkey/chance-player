// Package chance defines result value objects returned by use-cases.
// These are the data shapes that cross the use-case boundary into transport layers.
package chance

import "github.com/unmo/chance-player/internal/domain/shuffle"

// CoinFlipResult is the outcome of flipping a single coin.
type CoinFlipResult struct {
	Side shuffle.Side `json:"side"`
}

// MultiCoinFlipResult is the outcome of flipping multiple coins.
type MultiCoinFlipResult struct {
	Sides []shuffle.Side `json:"sides"`
	Count int            `json:"count"`
}

// DiceRollResult is the outcome of rolling a single die.
type DiceRollResult struct {
	Value int `json:"value"`
	Sides int `json:"sides"`
}

// MultiDiceRollResult is the outcome of rolling multiple dice.
type MultiDiceRollResult struct {
	Rolls []DiceRollResult `json:"rolls"`
	Sum   int              `json:"sum"`
	Sides int              `json:"sides"`
}

// CardDrawResult is the outcome of drawing a card.
type CardDrawResult struct {
	Suit  shuffle.Suit      `json:"suit"`
	Value shuffle.CardValue `json:"value"`
}

// RandomNumberResult is the outcome of generating a random number.
type RandomNumberResult struct {
	Number int `json:"number"`
	Min    int `json:"min"`
	Max    int `json:"max"`
}

// IChingLineType classifies a single I Ching line by its coin-sum value.
type IChingLineType string

const (
	// OldYin (sum=6): broken changing line — transforms to solid
	OldYin IChingLineType = "old_yin"
	// YoungYang (sum=7): solid stable line
	YoungYang IChingLineType = "young_yang"
	// YoungYin (sum=8): broken stable line
	YoungYin IChingLineType = "young_yin"
	// OldYang (sum=9): solid changing line — transforms to broken
	OldYang IChingLineType = "old_yang"
)

// IChingLineResult is the outcome of one line in a hexagram consultation.
type IChingLineResult struct {
	Coins    []shuffle.Side `json:"coins"`    // the 3 coin results
	Sum      int            `json:"sum"`       // 6, 7, 8, or 9
	LineType IChingLineType `json:"line_type"` // classification
}

// IChingConsultResult is the full 6-line hexagram consultation result.
type IChingConsultResult struct {
	Lines [6]IChingLineResult `json:"lines"`
}
