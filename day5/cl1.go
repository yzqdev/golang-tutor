package main

import (
	"fmt"
	"time"
)

func add(a int, b int, c chan int) {
	var sum int
	sum = a + b
	time.Sleep(3 * time.Second)
	c <- sum
}
func Ayy() {
	var str string
	str = "hello worlf"
	for i := 1; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
}
