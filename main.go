package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, newCard())

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
