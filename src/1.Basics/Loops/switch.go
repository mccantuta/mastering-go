package main

import (
	"fmt"
)

func main() {
	//switchIsPositive(1)
	//switchConditional()
	//switchWithDeclaration()
	//commaSeparatedValues(2)
	//commaSeparatedStringValues("3")
	verifyFallThrough(1)
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
	//x := 0
	//Is not necessary to include the x variable in the switch sentence, e.g. switch {
	switch x := 0; { //If we are including the declarated variables, we must to add semicolon.
	case x < 10:
		fmt.Printf("Less than 10: %d", x)
	case x >= 10:
		fmt.Println("Grather than 10: %d", x)
	}
}

func commaSeparatedValues(value int) {
	switch value {
	case 0, 2, 4, 6, 8:
		fmt.Printf("Odd: %d", value)
	case 1, 3, 5, 7, 9:
		fmt.Printf("Even: %d", value)
	}
}

func commaSeparatedStringValues(value string) {
	switch value {
	case "0", "2", "4", "6", "8":
		fmt.Printf("Odd: %v", value)
	case "1", "3", "5", "7", "9":
		fmt.Printf("Even: %v", value)
	}
}

func verifyFallThrough(value int) {
	//We don't need to add break sentence for each case clausule, because Go doesn't have automatic fall through
	switch value {
	case 0:
		fmt.Println("Is 0")
	case 1:
		fmt.Println("Is 1")
	case 2:
		fmt.Println("Is 2")
	}
}
