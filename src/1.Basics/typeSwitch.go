package main

import "fmt"

type Test struct {
}

func main() {
	getParameterType(Test{})
}

func getParameterType(parameter interface{}) {
	switch t := parameter.(type) {
	case bool:
		fmt.Printf("Boolean: %t", t)
	case int:
		fmt.Printf("Integer: %d", t)
	case string:
		fmt.Printf("String: %v", t)
	default:
		fmt.Printf("Unexpected type: %T", t)
	}
}
