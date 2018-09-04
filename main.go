package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p1 := person{firstName: "Mohammed", lastName: "Damia", contactInfo: contactInfo{email: "mdmaia@tesla.con", zipcode: 89408}}

	p1.updateName("Jacob")
	fmt.Println(p1)

	name := "Mohammed"
	fmt.Println(*&name)
}

func (p *person) updateName(name string) {
	p.firstName = name
	// fmt.Println(p)
}
