package main

import "fmt"

func main() {
	fmt.Println("Main function")
	a, b := twoResults()
	fmt.Println(a + b)
	a, c := twoResults() //variable a is redeclared, variable b is created
	fmt.Println(a + c)
	a, b := twoResults() // error, no new variables on left side of :=
	fmt.Println(a + c)
}

func twoResults() (int, int) {
	return 1, 2
}
