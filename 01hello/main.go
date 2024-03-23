package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Printf("%d \n", i)
	}
	end := time.Now()

	fmt.Printf("Time taken by Go's for loop: %v\n", end.Sub(start))
}
