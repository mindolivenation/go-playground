package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	total := len(cards)
	if total != 52 {
		t.Errorf("Excepted 52 but got %v", total)
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", cards[0])
	}

	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs but got %v", cards[len(cards)-1])
	}
}

func TestStoreDataAndRetrieveData(t *testing.T) {
	cards := newDeck()
	saveCardsToFile := cards.storeData("_storeDeckTesting.txt")
	loadFromFile := retrieveData("_storeDeckTesting.txt")
	if saveCardsToFile != nil {
		t.Errorf("Expected nil but got %v", saveCardsToFile)
	}

	if loadFromFile[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", cards[0])
	}

	if len(loadFromFile) != 52 {
		t.Errorf("Excepted 52 but got %v", len(loadFromFile))
	}

	os.Remove("_storeDeckTesting.txt")

}
