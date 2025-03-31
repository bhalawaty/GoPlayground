package main

import (
	"os"
	"testing"
)

func TestNewDeckCount(t *testing.T) {
	cards := newDeck()

	if len(cards) != 20 {
		t.Errorf("we expect the cards count is 20, we got %v", len(cards))
	}
}

func TestSaveDeckToFileAndReadFromFile(t *testing.T) {
	os.Remove("_testDeckFile")

	cards := newDeck()
	err := cards.saveToFile("_testDeckFile")
	if err != nil {
		t.Errorf("there is an issue with save to file name : %v", "_testDeckFile")
	}
	newCards := getNewDeckFromFile("_testDeckFile")

	if len(newCards) != 20 {
		t.Errorf("we expect the cards count is 20, we got %v", len(cards))
	}

	os.Remove("_testDeckFile")
}
