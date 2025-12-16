package main

func main() {

	//card := newCard()
	//cards := deck{"Ace of spades", newCard()}
	//cards = append(cards, "six of hearts")
	// fmt.Println(card)
	//fmt.Println(cards)
	//direct declaration
	//fmt.Println("using direct method")
	// for c := range cards {
	// 	fmt.Println(cards[c])
	// }
	// fmt.Println("using deck type")
	// cards.print()

	//cards := newDeck()
	// hand, remainingCards := cards.deal(5)
	// fmt.Println("hand")
	// hand.print()
	// fmt.Println("remiaining cards")
	// remainingCards.print()

	// fmt.Println("single string....")
	// fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shufffle()
	cards.print()

	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))

}

// func newCard() string {
// 	return "Ace of diamonds"
// }
