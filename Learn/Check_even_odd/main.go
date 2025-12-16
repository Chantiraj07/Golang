package main

import "fmt"

func main() {

	numberSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numberSlice {
		if number%2 == 0 {
			fmt.Println(number, ":: even")
		} else {
			fmt.Println(number, ":: odd")
		}
	}
}
