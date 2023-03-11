package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	alex := person{
		"Alex",
		"Anderson",
		contactInfo{
			"alex@test.io",
			666,
		},
	}

	max := person{
		firstname: "Max",
		lastname:  "Mustermann",
		contactInfo: contactInfo{
			email:   "max@test.io",
			zipCode: 12345,
		},
	}

	alex.updateName("Alexis")
	max.updateName("Maximilian")
	alex.print()
	max.print()
}

func (p *person) print() {
	fmt.Printf("%+v\n", *p)
}

func (p *person) updateName(newFirstName string) {
	p.firstname = newFirstName
}
