package main

import "fmt"

type Runner interface {
	run()
}

type Eater interface {
	eat()
}

type Mammal interface {
	Runner
	Eater
}

type Dog struct{}

func (d Dog) eat() {
	fmt.Println("Eating")
}

func main() {
	d := Dog{}
	d.eat()
}
