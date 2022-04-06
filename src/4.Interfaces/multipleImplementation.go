package main

import "fmt"

type interface1 interface {
	method1()
}

type interface2 interface {
	method1()
}

type S struct{}

func (s S) method1() {
	fmt.Println("Method 1")
}

func run1(i interface1) {
	fmt.Println("Run interface1")
}

func run2(i interface2) {
	fmt.Println("Run interface2")
}

func main() {
	s := S{}
	run1(s)
	run2(s)
}
