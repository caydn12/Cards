package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

//NewDeck : returns a initialized deck of cards
func NewDeck() deck {
	newDeck := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Deal : Returns a hand of cards, returns a new version of the deck without the dealt cards in it
func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//SaveToFile : Saves all cards in the deck to a file passed in by the user. Each card is seperated by a comma
func (d deck) SaveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//ReadFromFile : Reads in a deck's data from a file
func ReadFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	myDeck := deck{}
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	} else {
		compressedDeck := string(byteSlice)
		myDeck = deck(strings.Split(compressedDeck, ","))
	}

	return myDeck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
