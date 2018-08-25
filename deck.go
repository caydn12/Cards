package main

import "fmt"

type deck []string

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
