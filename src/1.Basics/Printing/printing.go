package main

import (
	"fmt"
	"os"
)

type Example struct {
	name   string
	age    int
	income float64
}

func main() {
	printFunctions()
	//printStructs()
}

func printFunctions() {
	//Format according to a format specifier and writes to standard output
	fmt.Printf("Hello %d\n", 10)

	//Formats using the default format for its operands and writes to Writer.
	//Spaces are added between operands when neither is a string.
	fmt.Fprint(os.Stdout, "Hello ", 10, "\n")

	//Formats using the default formats for its operands and writes to standard output.
	//Spaces are always added between operands and a newline is appended.
	fmt.Println("Hello", 10)
	fmt.Println("Hello", "this", "is", "multiple", "parameters")

	//Sprint formats using the default formats for its operands and returns the resulting string.
	//Spaces are added between operands when neither is a string.
	fmt.Println(fmt.Sprint("Hello ", 10))
}

func printStructs() {
	newExample := Example{"Test", 30, 100.50}

	//Print only values, the filed name will not be printed.
	fmt.Printf("%v\n", newExample)

	//It will print both field name and value
	fmt.Printf("%+v\n", newExample)

	//It will print the struct, also both field name and value
	fmt.Printf("%#v\n", newExample)

	//It will print the argument's type
	fmt.Printf("%T\n", newExample)
	fmt.Printf("%T\n", "stringText")
	fmt.Printf("%T\n", 10)
	fmt.Printf("%T\n", []int{1, 2})
}
