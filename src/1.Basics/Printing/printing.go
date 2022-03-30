package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello %d\n", 10)
	fmt.Fprint(os.Stdout, "Hello ", 10, "\n")
	fmt.Println("Hello", 10)
	fmt.Println(fmt.Sprint("Hello ", 10))
}
