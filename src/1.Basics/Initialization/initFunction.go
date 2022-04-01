package main

import "fmt"

//Allowed to have more than one init function
func init() {
	fmt.Println("In Init function")
}

func init() {
	fmt.Println("In second Init function")
}

func main() {

}
