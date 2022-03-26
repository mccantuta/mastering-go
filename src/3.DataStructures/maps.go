package main

import "fmt"

func main() {
	//keyNotFound()
	commaOk()
}

func initialization() {
	var myMap = map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}
	for k, v := range myMap {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(myMap["dos"])

}

func keyNotFound() {
	//If the key is not found, map wil return the 'zeroed' value
	var myMap = map[int]string{
		1: "uno",
		2: "dos",
	}
	fmt.Println(myMap[3])
}

func commaOk() {
	var myMap = map[int]string{
		1: "uno",
		2: "dos",
	}
	v, ok := myMap[3]
	fmt.Println(v)
	fmt.Println(ok)
}
