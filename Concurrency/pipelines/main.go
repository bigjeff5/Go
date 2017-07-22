package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline
	c := gen(1, 2, 3, 4, 5, 6, 7, 8, 9)
	out := sq(c)

	// Consume the output.

	for prnt := range out {
		fmt.Println(prnt) // 4
	}
	// fmt.Println(<-out) // 1
	// fmt.Println(<-out) // 9
	// fmt.Println(<-out) // 25

}
