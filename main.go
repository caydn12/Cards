package main

func main() {
	myDeck := deck{}

	myDeck = ReadFromFile("myDeck")

	myDeck.Print()
}
