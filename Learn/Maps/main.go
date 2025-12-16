package main

import "fmt"

func main() {
	colors := map[string]string{ // one way to create map
		"red":   "#ff0000",
		"green": "#qwer123",
		"white": "#ffffff",
	}

	//var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff" // update or add keys

	// delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
