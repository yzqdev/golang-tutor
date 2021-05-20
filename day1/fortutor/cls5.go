package main

import (
	"fmt"
	"time"
)

func main() {
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
