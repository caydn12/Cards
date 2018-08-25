package main

func main() {
	myDeck := NewDeck()

	for i := 0; i < 100; i++ {
		myDeck.Shuffle()
	}

	myDeck.Print()
}
