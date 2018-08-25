package main

import (
	"os"
	"testing"
)

func TestNewDeckForCorrectLength(t *testing.T) {
	deck := NewDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52. Got: %v", len(deck))
	}
}

func TestNewDeckForNoDuplicates(t *testing.T) {
	deck := NewDeck()

	duplicates := 0
	for _, card := range deck {
		duplicates = 0
		for _, compareCard := range deck {
			if card == compareCard {
				duplicates++
			}
		}
		if duplicates > 1 {
			t.Errorf("Duplicate cards have been found: %v", card)
		}
	}
}

func TestSaveToFileForErrorAndTestReadFromFileForCorrectLength(t *testing.T) {
	os.Remove("_decktesting")

	deck := NewDeck()

	err := deck.SaveToFile("_decktesting")

	if err != nil {
		t.Errorf("Error occured during the SaveToFile Call: %v", err)
	}

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52. Got: %v", len(deck))
	}

	os.Remove("_decktesting")
}
