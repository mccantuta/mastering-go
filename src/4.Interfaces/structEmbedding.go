package main

import "fmt"

type Address struct {
	street string
	number int
}

type Person struct {
	name string
	Address
}

func (a Address) fullAddress() {
	fmt.Println(a.street, a.number)
}

func main() {
	a := Address{"Street1", 123}
	a.fullAddress()
	p := Person{"Name", a}
	p.fullAddress()
}
