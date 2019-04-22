package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Diamond, Ace})
	fmt.Println(Card{Club, King})
	fmt.Println(Card{Spade, Three})
	fmt.Println(Card{Heart, Queen})
	// Output:
	// Ace of Diamonds
	// King of Clubs
	// Three of Spades
	// Queen of Hearts
}

func TestNew(t *testing.T) {
	deck := New()
	if deck.Size() != 52 {
		t.Errorf("deck does not have the expected count of cards: got %d instead of 52", deck.Size())
	}
}

func TestTake(t *testing.T) {
	deck := New()
	initialSize := deck.Size()

	card := deck.Take()
	if deck.Size() != initialSize-1 {
		t.Errorf("deck Size did not decrement by 1: got %d instead of %d", deck.Size(), initialSize-1)
	}

	if card.Value != Ace && card.Suit != Spade {
		t.Errorf("expected the card taken to be an Ace of Spades, got %s", card.String())
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(Sort(DefaultLess))
	expected := Card{Suit: Spade, Value: Ace}
	got := deck.Cards[0]
	if got != expected {
		t.Errorf("first card: expected %s, got %s", expected.String(), got.String())
	}

	expected = Card{Suit: Heart, Value: King}
	got = deck.Cards[deck.Size()-1]
	if got != expected {
		t.Errorf("last card: expected %s, got %s", expected.String(), got.String())
	}
}

func TestCustomSort(t *testing.T) {
	deck := New(Sort(customLess))
	expected := Card{Suit: Club, Value: Ace}
	got := deck.Cards[0]
	if got != expected {
		t.Errorf("first card: expected %s, got %s", expected.String(), got.String())
	}

	expected = Card{Suit: Spade, Value: Two}
	got = deck.Cards[deck.Size()-1]
	if got != expected {
		t.Errorf("last card: expected %s, got %s", expected.String(), got.String())
	}
}

func TestShuffle(t *testing.T) {
	deck1 := New(Shuffle)
	deck2 := New(Shuffle)

	if sameDeck(deck1, deck2) {
		t.Errorf("two shuffled decks must be different")
	}
}

func TestWithJokers(t *testing.T) {
	deck := New(WithJokers(5))

	var count int
	for _, c := range deck.Cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 5 {
		t.Errorf("number of Jokers did not match: expected %d, got %d", 5, count)
	}
}

func TestFilterCards(t *testing.T) {
	deck := New(Filter([]Card{Card{Suit: Heart}, Card{Value: Three}, Card{Suit: Club, Value: Five}}))

	for _, c := range deck.Cards {
		if c.Suit == Heart {
			t.Errorf("unexpected suit found: %s", c.Suit)
		}
		if c.Value == Three {
			t.Errorf("unexpected value found: %s", c.Value)
		}
		if c.Suit == Club && c.Value == Five {
			t.Errorf("unexpected card found: %s", c)
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	deck := New(Multiple(3))

	if deck.Size() != 52*3 {
		t.Errorf("wrong number of cards in multiple deck: expected %d, got %d", 52*3, deck.Size())
	}
}

func TestMixedOptions(t *testing.T) {
	deck := New(WithJokers(4), Multiple(2), Filter([]Card{Card{Value: Two}}), Sort(customLess))
	expectedSize := ((52 + 4) * 2) - 8

	if deck.Size() != expectedSize {
		t.Errorf("wrong number of cards: expected %d, got %d", expectedSize, deck.Size())
	}
}

func customLess(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].String() < cards[j].String()
	}
}

func sameDeck(d1, d2 *Deck) bool {
	for i := range d1.Cards {
		if d1.Cards[i] != d2.Cards[i] {
			return false
		}
	}
	return true
}
