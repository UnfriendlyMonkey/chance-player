package shuffle

// Suit represents a card suit.
type Suit string

const (
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Spades   Suit = "spades"
	Clubs    Suit = "clubs"
)

// CardValue represents a card face value.
type CardValue string

var cardValues = []CardValue{"6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var cardSuits = []Suit{Hearts, Diamonds, Spades, Clubs}

// Card is a playing card with a suit and value.
type Card struct {
	Suit  Suit      `json:"suit"`
	Value CardValue `json:"value"`
}

// CardDeck holds a set of cards and can draw one at random.
type CardDeck struct {
	cards []Card
}

// NewStandard36Deck returns a 36-card deck (4 suits × 9 values: 6 through Ace).
func NewStandard36Deck() CardDeck {
	cards := make([]Card, 0, 36)
	for _, s := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, Card{Suit: s, Value: v})
		}
	}
	return CardDeck{cards: cards}
}

// Pick returns a random card from the deck.
func (d CardDeck) Pick(src RandomSource) Card {
	return d.cards[src.Intn(len(d.cards))]
}
