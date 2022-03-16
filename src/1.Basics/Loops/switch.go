package main

import (
	"fmt"
)

func main() {
	//switchIsPositive(1)
	//switchConditional()
	switchWithDeclaration()
}

func switchIsPositive(x int) {
	//switch without variable declaration
	switch {
	case x < 0:
		fmt.Println("Negative")
	case x >= 0:
		fmt.Println("Positive")
	}
}

func switchConditional() {
	switch condition := 5; {
	case condition < 10:
		fmt.Println("Less than 10")
	case condition >= 10:
		fmt.Println("Grather than 10")
	}
}

func switchWithDeclaration() {
	//Is not necessary to include the x variable in the switch sentence, e.g. switch {
	switch x;{ //If we are including the declarated variables, we must to add semicolon.
	case x < 10:
		fmt.Println("Less than 10")
	case x >= 10:
		fmt.Println("Grather than 10")
	}
}

func commaSeparatedValues(int value) {
	switch value; {
	case 0, 2, 4, 6, 8:
		fmt.Println("Odd")
	case 1, 3, 5, 7, 9:
		fmt.Println("Even")
	}

