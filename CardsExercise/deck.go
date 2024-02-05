package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// add a new function for this type to print out each individual card in this deck

// this function has a different structure because it has a receiver, we call the (d deck) part a receiver
// this gives us the cards.print() method call capability
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this function doesnt need a receiver since it doesnt have any input
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two", "Three","Four"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// deal => create a hand of cards
// the second paranthesis is the return types, since this method returns 2 items
// why we are not using receiver here? If I call cards.deal(5), it sounds like I will change cards, but that's not the case here
// we are not changing card we are just returning a deal of 5 from the cards
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:] // This will give us a slice of the deck from 0 to the handsize, and then return the rest
	// of the remaining cards from handsize till the end of the list
}

// We use receiver because we want to be able to call cards.ToString()
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666);
}

func newDeckFromFile(filename string) deck {
	byteSlice, error := os.ReadFile(filename)
	if error != nil {
		// log the error and entirely quit the program
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")
	return stringSlice
}

func (d deck) shuffle() {
	// This 2 lines is to generate a new random sequence each time by generting different seeds for the random generator
	source := rand.NewSource(time.Now().UnixNano()) // create new source
	r := rand.New(source) // create a new Rand type with the generate source

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i] , d[newPosition] = d[newPosition] , d[i]
	}
}