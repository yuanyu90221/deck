package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	firstCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != firstCard {
		t.Error("Expected  Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	firstCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != firstCard {
		t.Error("Expected  Ace of Spades as first card. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jockers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, received:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expeced %d cards, received %d cards", 13*4*3, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffle  deterministic
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shutfle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %s, received %s", second, cards[1])
	}
}
