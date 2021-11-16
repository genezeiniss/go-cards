package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type 'deck' that extends slide of strings
// array - fixed length list
// slide - an array that can graw or shrink
type deck [] string

func newDeck() deck {

	cards := deck{} // creates an empty slice of strings

	cardSuits := [] string { "Spades", "Diamonds", "Hearts", "Clubs" }
	cardValues := [] string { "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Knave", "Queen", "King", "Ace" }

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suite)
		}
	}

	return cards
}

// the parameter playingCards of deck type is a receiver of function print()
func (d deck) print() {

	for index, card := range d {
		fmt.Println(index, card)
	}
}

// return multiple values values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {

	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil { // 'nil' represents no error
		fmt.Println("error!" , err)
		os.Exit(1) // Conventionally, code zero indicates success, non-zero an error.
	}

	return strings.Split(string(byteSlice), ",")
}

func (d deck) shuffle()  {
	for index := range d {
		newPosition := rand.Intn(len(d) - 1)
		// swap the elements
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}