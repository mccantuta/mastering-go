package main

import "math"

const (
	ONE = 1
	// Fail because the function call to math.Sin needs to happen at run time
	CALCULATE = math.Sin(math.Pi / 4)
)

func main() {

}
