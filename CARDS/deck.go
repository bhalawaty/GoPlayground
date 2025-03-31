package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsValues := []string{"One", "Two", "Three", "Four", "Five"}
	cardsShapes := []string{"Ace", "Diamonds", "Hearts", "Clubs"}

	for _, shape := range cardsShapes {
		for _, value := range cardsValues {
			cards = append(cards, value+" OF "+shape)
		}
	}
	return cards
}
func deal(deck deck, handSize int) (deck, deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck deck) toString() string {
	return strings.Join(deck, ",")
}

func (deck deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(deck.toString()), 0666)
}

func getNewDeckFromFile(fileName string) []string {
	cards, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("error :", error)
		os.Exit(500)
	}

	cardsStr := strings.Split(string(cards), ",")
	return cardsStr

}
func (deck deck) print() {
	for index, card := range deck {
		fmt.Println(index, card)
	}
}

func (deck deck) shuffle() {
	for index := range deck {
		// we can create a random by this way too
		source := rand.NewSource(time.Now().UnixNano())
		ra := rand.New(source)
		newRandomNum := ra.Intn(len(deck) - 1)
		deck[index], deck[newRandomNum] = deck[newRandomNum], deck[index]
	}
}
