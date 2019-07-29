package main

func main() {
	// cards := newDeck()
	// cards.saveToFile(("myCards"))
	cards := newDeckFromFile("myCards")
	cards.print()
}
