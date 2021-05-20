package main

import (
	"fmt"
	"math"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Print("number", math.Sqrt(8))
	fmt.Println(split(17))
}
