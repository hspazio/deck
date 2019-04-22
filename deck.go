package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Deck struct {
	Cards []Card
}

// Size returns how many cards are in the deck
func (d *Deck) Size() int {
	return len(d.Cards)
}

// Take removes and returns the first card on the top of the deck
func (d *Deck) Take() Card {
	card := d.Cards[0]
	d.Cards[0] = d.Cards[d.Size()-1]
	d.Cards[d.Size()-1] = Card{}
	d.Cards = d.Cards[:d.Size()-1]
	return card
}

func (d *Deck) add(c ...Card) {
	d.Cards = append(d.Cards, c...)
}

// New creates a deck of cards and takes in input a list of options to
// customize the deck
func New(opts ...func(*Deck)) *Deck {
	deck := &Deck{}
	for _, suit := range suits {
		for value := Ace; value <= King; value++ {
			deck.add(Card{Suit: suit, Value: value})
		}
	}
	for _, opt := range opts {
		opt(deck)
	}
	return deck
}

// CardInList returns true if a group of cards contains the card for the specified value
func CardInList(c Card, list []Card) bool {
	for _, item := range list {
		if c.Suit == item.Suit && item.Value == 0 || // filter specifies only Suit
			c.Value == item.Value && item.Suit == 0 || // filter specifies only Value
			c.Value == item.Value && c.Suit == item.Suit { // filter specifies both
			return true
		}
	}
	return false
}

// Sort is an option to customize how we want the cards to be ordered
func Sort(less func(cards []Card) func(i, j int) bool) func(*Deck) {
	return func(d *Deck) {
		sort.Slice(d.Cards, less(d.Cards))
	}
}

// DefaultLess is a basic comparison function that can be used with Sort option
func DefaultLess(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].absValue() < cards[j].absValue()
	}
}

// WithJokers adds a number of Jokers (special cards) to the deck
func WithJokers(n int) func(*Deck) {
	return func(d *Deck) {
		for i := 0; i < n; i++ {
			d.add(Card{Suit: Joker})
		}
	}
}

// Filter is an option to remove cards from the deck that match certain criteria
func Filter(toFilter []Card) func(*Deck) {
	return func(d *Deck) {
		filtered := d.Cards[:0]
		for _, c := range d.Cards {
			if !CardInList(c, toFilter) {
				filtered = append(filtered, c)
			}
		}
		d.Cards = filtered
	}
}

// Multiple is an option that defines how many decks we want to play with
func Multiple(n int) func(*Deck) {
	return func(d *Deck) {
		cards := d.Cards
		for i := 0; i < n-1; i++ {
			d.add(cards...)
		}
	}
}

// Shuffle is an option to sort the deck in random order
func Shuffle(d *Deck) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
