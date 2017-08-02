package main

import (
	"fmt"
)

func main() {
	// sets the largest value target the program can process
	target_limit := 999

	coin_value := []int{1, 5, 10, 25}
	target_value := 100

	min_coins := make([]int, target_value+1)

	num_coins := len(coin_value)

	//initialize each slot in min_coins to target_limit + 1, leave 0 at 0
	for i := 1; i < len(min_coins); i++ {
		min_coins[i] = target_limit + 1
	}

	for sub_value := 1; sub_value <= target_value; sub_value++ {
		for j := 0; j < num_coins; j++ {
			Vj := coin_value[j]
			if (Vj <= sub_value) && (min_coins[sub_value-Vj]+1 < min_coins[sub_value]) {
				min_coins[sub_value] = min_coins[sub_value-Vj] + 1
			}
		}
	}
	fmt.Println("Sum	:    Min # of coins")
	for i := 0; i < target_value; i++ {
		fmt.Println(i, "	:	", min_coins[i])
	}

	fmt.Println(target_value, "	:	", min_coins[target_value])

}
