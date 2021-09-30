package main

func main() {
	cards := newDeck()
	cards.shuffle()
	deal(cards, 5)
	cards.print()

}
