package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()

	}()
	wg.Wait()
	fmt.Println("Main func completed")
}

func work() {
	time.Sleep(2 * time.Second)
	fmt.Println("Finishing work method")
}
