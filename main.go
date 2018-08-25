package main

import "fmt"

func main() {
	cards := deck{}
	cards = append(cards, newCard())

	cards.Print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
