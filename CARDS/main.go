package main

func main() {
	//cards := newDeck()
	//hand, remainingCards := deal(cards, 5)

	//cards.saveToFile("bilalTestFile")
	//hand.print()
	//remainingCards.print()

	//cards := getNewDeckFromFile("bilalTestFile")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
