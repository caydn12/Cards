package main

import "fmt"

type deck []string

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
