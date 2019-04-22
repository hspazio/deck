package deck

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Diamond, Ace})
	fmt.Println(Card{Club, King})
	fmt.Println(Card{Spade, Three})
	fmt.Println(Card{Heart, Queen})
	fmt.Println(Card{Joker, 0})
	// Output:
	// Ace of Diamonds
	// King of Clubs
	// Three of Spades
	// Queen of Hearts
	// Joker
}
