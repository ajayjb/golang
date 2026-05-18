package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName  string
	secondName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) changeFirstName(firstName string) {
	// a := &p
	// fmt.Printf("%p\n", a)
	(*p).firstName = firstName
}

func main() {

	var newUser person

	newUser.firstName = "Ajay"
	newUser.secondName = "J B"

	newUser.contactInfo = contactInfo{
		email:   "ajayjb11@gmail.com",
		zipCode: 49494,
	}

	// a := &newUser
	// fmt.Printf("%p\n", a)

	newUser.changeFirstName("Jackie")

	newPersonPointer := &newUser

	(*newPersonPointer).firstName = "lokie"

	// newUser := person{
	// 	firstName:  "Ajay",
	// 	secondName: "J B",
	// 	contactInfo: contactInfo{
	// 		email:   "ajayjb11@gmail.com",
	// 		zipCode: 49494,
	// 	},
	// }

	newUser.print()

}
