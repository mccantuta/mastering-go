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

	fmt.Println(stringValue)
}
