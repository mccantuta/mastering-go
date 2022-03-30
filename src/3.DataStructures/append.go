package main

import "fmt"

func main() {
	//appendSingleElement()
	//appendMultipleElements()
	appendSliceToSlice()
}

func appendSingleElement() {
	slice1 := []int{1, 2}
	slice1 = append(slice1, 3)
	fmt.Println(slice1)
}

func appendMultipleElements() {
	slice1 := []int{1, 2}
	slice1 = append(slice1, 3, 4, 5)
	fmt.Println(slice1)
}

func appendSliceToSlice() {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}
