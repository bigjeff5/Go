package main

import (
	"fmt"

	"github.com/polaris/bitmap"
	// Import "bitmap" from polaris github
)

func main() {
	size := 100000000000
	bmap := bitmap.NewBitmapSize(size)

	fmt.Println(len(bmap))
}
