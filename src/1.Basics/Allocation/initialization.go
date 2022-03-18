package main

type Person struct {
	name    string
	address string
}

func main() {
	NewPerson()
	NewPersonNamedFields()
	initializationWithNew()
}

func NewPerson() *Person {
	person := Person{"Name", "Address"}
	return &person
}

func NewPersonNamedFields() *Person {
	person := Person{address: "Address", name: "Name"}
	return &person
}
