package main

import "fmt"

func main() {
	forWithInit()
	forLikeWhile(1)
	forInfinite(1)
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
