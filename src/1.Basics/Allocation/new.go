package main

import "fmt"

type Person struct {
	name string
}

func initializeAsPointer() {
	p := new(Person) // type *Person
	fmt.Println(p)
	fmt.Println(p.name)
}

func initializeAsVariable() {
	var p Person // type Person
	fmt.Println(p)
	fmt.Println(p.name)
}

func main() {
	initializeAsPointer()
	initializeAsVariable()
}
