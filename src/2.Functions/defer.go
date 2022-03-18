package main

import "fmt"

func main() {
	//stackDefer()
	deferedFunction()
	panic("Error")
}

func deferedFunction() {
	//Defered funcion runs also when there is a panic
	fmt.Println("I was defered")
}

func stackDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
