package main

func main() {
	// var card string = "Ace of Spades"
	//fmt.Println([]byte("Hi there")) //type conversion

	card := "Ace of Spades" //first declaration

	cards := deck{card, card}
	cards = newDeck()

	cards.shuffle()
	cards.print()

	//deal the cards with a hand of 5
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// fmt.Println()
	// remainingDeck.print()

	//fmt.Println(cards.toString())

	// cards.saveToFile("testWriteToFile.tmp")
	// newDeck := newDeckFromFile("testWriteToFile.doc")
	// newDeck := newDeckFromFile("testWriteToFile.tmp")
	// newDeck.print()

}
