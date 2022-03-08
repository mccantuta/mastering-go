package main

import "fmt"

func main() {
	//forWithInit()
	//forLikeWhile(1)
	//forInfinite(1)
	//forWithSlice()
	//forWithSliceOnlyPosition()
	//forWithSliceOnlyValue()
	forWithStrings()
}

func forWithInit() {
	// for init; condition; post {}
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func forLikeWhile(i int) {
	// for condition {}
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func forInfinite(i int) {
	// for {}
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}
}

func forWithSlice() {
	list := []int{1, 2, 3}
	for position, value := range list {
		fmt.Println(position)
		fmt.Println(value)
	}
}

func forWithSliceOnlyPosition() {
	list := []int{1, 2, 3}
	for position := range list {
		fmt.Println(position)
	}
}

func forWithSliceOnlyValue() {
	list := []int{1, 2, 3}
	for _, value := range list {
		fmt.Println(value)
	}
}

func forWithStrings() {
	cad := "This is a text line"
	for position, value := range cad {
		fmt.Printf("The character %#U is in position %d\n", value, position)
	}
}
