package main

import "fmt"

func main() {

	card := newCard()
	cards := []string{"Ace of spades", newCard()}
	cards = append(cards, "six of hearts")

	fmt.Println(card)
	fmt.Println(cards)

	for c := range cards {
		fmt.Println(cards[c])
	}
}

func newCard() string {
	return "Ace of diamonds"
}
