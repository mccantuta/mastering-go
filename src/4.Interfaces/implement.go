package main

import "fmt"

type Animal interface {
	breathe()
	eat()
	mood()
}

type Dog struct{}

type Cat struct{}

func (d Dog) breathe() {
	fmt.Println("Dog Breathing")
}

func (d Dog) eat() {
	fmt.Println("Dog Eating")
}

func (d Dog) mood() {
	fmt.Println("Happy dog")
}

func (c Cat) breathe() {
	fmt.Println("Cat Breathing")
}

func (c Cat) eat() {
	fmt.Println("Cat Eating")
}

func (c Cat) mood() {
	fmt.Println("Happy cat")
}

func usingStructs() Dog {
	var d Dog = Dog{}
	d.breathe()
	return d
}

func usingInterface() Animal {
	var a Animal
	a = Dog{}
	a.breathe()
	return a
}

func toAdopt(a Animal) {
	a.mood()
}

func main() {
	//d := usingStructs()
	//a := usingInterface()

	c := Cat{}
	//d := Dog{}

	toAdopt(c)
}
