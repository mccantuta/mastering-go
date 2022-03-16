package main

import "fmt"

func main() {
	//key, value := result()
	key, value := resultAsVariables()
	fmt.Println(key)
	fmt.Println(value)
}

func result() (key string, value string) {
	return "1", "2"
}

func resultAsVariables() (key string, value string) {
	key = "1"
	value = "2"
	return
}
