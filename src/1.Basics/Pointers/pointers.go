package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{"Name", 100}
	update(&person)
	fmt.Println(person)

	person = Person{"Name", 100}
	nonUpdate(person)
	fmt.Println(person)
}

func update(person *Person) {
	person.name = "Updated"
}

func nonUpdate(person Person) {
	person.name = "Non updated"
}
