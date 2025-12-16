package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("Interfacesss")
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
func (eb englishBot) getGreeting() string {
	//After writing some code to generate english greeting
	return "Hello there"
}
func (sb spanishBot) getGreeting() string {
	////After writing some code to generate spanish greeting
	return "Hola!!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// no overloading available

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
