//go:generate stringer -type=Value,Suit

package deck

import "fmt"

// Suit represents is the figure of the card
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // special suit
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Value is the instrinsic value of the card
type Value uint8

const (
	_ Value = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card is the combination of Suit and Value
type Card struct {
	Suit  Suit
	Value Value
}

// String describes the card
func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Value.String(), c.Suit.String())
}

func (c Card) absValue() int {
	return int(c.Suit)*int(King) + int(c.Value)
}
