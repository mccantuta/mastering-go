package main

import "fmt"

func main() {
	//initalizeSlice()
	//basicFunction()
	relatedSlice()
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
	fmt.Println(slice0)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

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

func relatedSlice() {
	slice1 := []int{1, 2}
	slice2 := append(slice1, 3)
	// Due to slice's capacity was 2, a new slice was created,
	//the append method return a reference value
	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice3 := append(slice2, 4)
	// Due to slice's capacity is 4, the underline array remains,
	// the append method return a copy value
	slice2[0] = 200
	fmt.Println(slice2)
	fmt.Println(slice3)
}
