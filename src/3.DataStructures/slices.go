package main

import "fmt"

func main() {
	initalizeSlice()
}

func initializeSlice() {
	var slice0 []int
	var slice1 = []int{1, 2, 3}
	var slice2 = make([]int, 4, 4)

	//When the capacity argument is omitted, it defaults to the specified length
	var slice3 = make([]int, 4)

	//When the length argument is smaller than capacity, the last elements will be `null`
	var slice4 = make([]int, 2, 4)
	//Now it's [0 0 nil nil]
}

func basicFunction() {
	mySlice := make([]int, 3)
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3

	for i := 0; i < 3; i++ {
		fmt.Println(mySlice[i])
	}
}
