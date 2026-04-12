package shuffle

// Side represents one face of a coin.
type Side string

const (
	Heads Side = "heads"
	Tails Side = "tails"
)

// Coin is a two-sided randomizable item.
type Coin struct{}

var coinSides = []Side{Heads, Tails}

// Pick returns a random coin side using the provided RandomSource.
func (c Coin) Pick(src RandomSource) Side {
	return coinSides[src.Intn(len(coinSides))]
}
