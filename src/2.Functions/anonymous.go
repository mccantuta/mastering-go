package main

import "fmt"

func main() {
	//anonymuousFunction()
	returnFunction()("Message to returned function")
}

func anonymuousFunction() {
	func(message string) {
		fmt.Println(message)
	}("Anonymous function")
}

func returnFunction() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}
