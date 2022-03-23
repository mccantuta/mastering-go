package main

import "fmt"

func main() {
	//declaration [capacity], index starts at zero
	myArray := [3]int{}
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	for i := 0; i < 3; i++ {
		fmt.Println(myArray[i])
	}
}

func initializeArray() {
	//Will be filled with zeroed values
	var array1 = [3]int{}

	//Number of values between {} can not be larger than size
	var array2 = [3]int{1, 2, 3}

	//The compiler will count the array elements for you
	var array3 = [...]int{1, 2, 3}
}
