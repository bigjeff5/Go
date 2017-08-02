package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boxMover(s *int, d *int, mover int) {
	val := *s
	*d = val
	*s = 0

	fmt.Println("Box Mover", mover, "is moving a box!")
	
	rand.Seed(rand.Int63())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))

	fmt.Println("Box Mover", mover, "has moved his box!")
}

func main() {

	boxShelf := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	boxIncinerator := []int{99, 99, 99, 99, 99, 99, 99, 99, 99, 99}

	for i := 0; i < 10; i++ {
		go boxMover(&boxShelf[i], &boxIncinerator[i], i)
	}

	time.Sleep(time.Millisecond * 100)

	fmt.Println("Boxes in Incinerator:")
	for i := range boxIncinerator {
		fmt.Println(boxIncinerator[i])
	}
}
