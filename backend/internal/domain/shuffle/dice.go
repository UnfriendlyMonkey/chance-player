package shuffle

// Dice is a numbered die with configurable sides.
type Dice struct {
	sides int
}

// NewDice creates a die with the given number of sides (e.g. 6 for a standard d6).
// Faces are numbered 1 through sides inclusive.
func NewDice(sides int) Dice {
	if sides < 2 {
		sides = 6
	}
	return Dice{sides: sides}
}

// Sides returns the number of faces on the die.
func (d Dice) Sides() int { return d.sides }

// Pick returns a random face value in [1, sides].
func (d Dice) Pick(src RandomSource) int {
	return src.Intn(d.sides) + 1
}
