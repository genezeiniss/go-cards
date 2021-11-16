package main

import (
	"os"
	"reflect"
	"testing"
)

// test functions name are started from capital test word, such as TestFunc(), and should include a tested function name
// an opposite of application functions that started by lower case letter.

func TestNewDeck(t *testing.T)  {

	deck := newDeck()

	expectedDeckLength := 52
	if len(deck) != expectedDeckLength {
		t.Errorf("deck length assertion failed. \n" +
			"expected: %v \n" +
			"actual: %v", expectedDeckLength, len(deck))
	}

	expectedFirstCard := "Two of Spades"
	if deck[0] != expectedFirstCard {
		t.Errorf("first card in deck assertion failed. \n" +
			"expected: %v \n" +
			"actual: %v", expectedFirstCard, deck[0])
	}

	expectedLastCard := "Ace of Clubs"
	if deck[len(deck) - 1] != expectedLastCard {
		t.Errorf("last card in deck assertion failed. \n" +
			"expected: %v \n" +
			"actual: %v", expectedLastCard, deck[len(deck) - 1])
	}
}

//goland:noinspection ALL
func TestSaveToFileAndNewDeckFromFile(t *testing.T)  {

	fileName := "_deck_testing"

	// remove all dirty or leftover files
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)
	expectedDeckLength := 52
	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("deck from file length assertion failed. \n" +
			"expected: %v \n" +
			"actual: %v", expectedDeckLength, len(loadedDeck))
	}

	// finally, remove created file
	os.Remove(fileName)
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	handSize := 6
	cardsInHand, remainingCards := deal(deck, handSize)

	if len(cardsInHand) != handSize {
		t.Errorf("deal function error: cards in hand assertion failed \n" +
			"expected: %v \n" +
			"actual: %v", handSize, len(cardsInHand))
	}

	expectedRemainingCards := 52 - handSize
	if len(remainingCards) != expectedRemainingCards {
		t.Errorf("deal function error: remining cards assertion failed \n" +
			"expected: %v \n" +
			"actual: %v", expectedRemainingCards, len(remainingCards))
	}
}

func TestShuffle(t *testing.T) {
	notShuffledDeck := newDeck()
	deck := newDeck()
	deck.shuffle()

	if reflect.DeepEqual(notShuffledDeck, deck) {
		t.Errorf("the deck wasn't shuffled as expected")
	}
}