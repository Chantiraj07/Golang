package main

import "fmt"

//Normal struct
// type person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// func main() {
// 	fmt.Println("Hello structs")
// 	alex := person{"alex", "anderson", 98}
// 	fmt.Println(alex)
// 	andrew := person{firstName: "andrew", lastName: "baker"}
// 	fmt.Println(andrew)
// 	alex.firstName = "not alex"
// 	fmt.Println(alex)
// 	fmt.Printf("%+v", alex)
// }

//Multi struct
// type contactInfo struct {
// 	email string
// 	zip   int
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   []contactInfo

// 	//contactInfo contactInfo or contactInfo same same
// }

// func main() {
// 	fmt.Println("Hii!! there")
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Party",
// 	}
// 	jim.contact = append(jim.contact, contactInfo{
// 		email: "abc@gmail.com",
// 		zip:   12345767,
// 	})
// 	fmt.Printf("%+v", jim)
// 	fmt.Println("\n Using receiver function")
// 	jim.showInfo()
// }
// func (p person) showInfo() {
// 	fmt.Printf("%+v", p)
// }

//Pointers

type person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("Pointerssss")

	jim := person{firstName: "jim", lastName: "Partyy"}
	fmt.Printf("%+v", jim)
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	//same as above

	jim.updateName("jimmy")
	jim.showInfo()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) showInfo() {
	fmt.Printf("%+v", p)
}
