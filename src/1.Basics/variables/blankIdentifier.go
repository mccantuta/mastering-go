package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	//Error because there is not any new variable to assign with `:=` operator
	for _, _ := range slice {
		fmt.Println("range")
	}
}
