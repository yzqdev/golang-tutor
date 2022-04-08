package calc

import (
	"fmt"
	"math"
	"reflect"
)

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func BitOperation() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func Add(a int, b int) int {
	return a + b
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func simplePrint() {
	var i int
	fmt.Println(i, c, java, python)
	fmt.Print("number", math.Sqrt(8))
	fmt.Println(split(17))
}
func getType() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(f), reflect.TypeOf(z))
}
