package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(f), reflect.TypeOf(z))
}
