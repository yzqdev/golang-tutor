package calc

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func ForPrint() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Print(sum)
}
func ForPrint2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	//go语言没while语句
	fmt.Println(sum)
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func PowPrint() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
func SwitchPrint() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func SwitchPrint2() {
	fmt.Println("when is staturay")
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(time.Now().Hour())
	switch time.Saturday {
	case today + 0:
		fmt.Println("tody")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
