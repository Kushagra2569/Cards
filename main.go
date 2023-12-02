package main

func main() {
	cards := newDeck()

	var handSize int = 5
	hand, remainingCards := deal(cards, handSize)

	hand.print()
	remainingCards.print()
}
