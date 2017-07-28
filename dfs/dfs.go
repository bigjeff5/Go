package main

import (
	"fmt"
)

func main() {
	size := 50000
	grid := make([][]byte, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]byte, size)
	}

	dfs(0, 0, grid)
/*
	for x := range grid {
		fmt.Println(grid[x])
	}
*/
	totalObjects := size * size
	fmt.Printf("The value at 0,0 is: %v\n", grid[0][0])
	fmt.Printf("The value at %d, %d is: %v\n", size-1, size-1, grid[size-1][size-1])
	fmt.Printf("The total number of values set is: %v\n", totalObjects)
}

func dfs(x1, y1 int, grid [][]byte) {
	grid[x1][y1] = 1
	y2 := y1 + 1
	if y1 < len(grid[x1])-1 && grid[x1][y2] != 1 {
		dfs(x1, y2, grid)
	}
	x2 := x1 + 1
	if x1 < len(grid)-1 && grid[x2][y1] != 1 {
		dfs(x2, y1, grid)
	}
}
