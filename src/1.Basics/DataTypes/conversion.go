package main

import "fmt"

func main() {
	convertToString("10")
}

func convertToString(value interface{}) {
	stringValue, ok := value.(string)
	if !ok {
		fmt.Println("Value is not a String")
	}
	// Other alternative is to use a one-line sentence like
	// if value, ok := value.(string); !ok {}

	fmt.Println(stringValue)
}
