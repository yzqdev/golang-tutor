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
func ayy() {
	var str string
	str = "hello worlf"
	for i := 1; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
}
func main() {
	//sum,avg:=Calculate(100,400)
	//fmt.Println(sum,avg)
	ayy()
	//var pipe chan int
	//pipe = make(chan int, 3)
	//go add(2, 5, pipe)
	//sum := <-pipe
	//fmt.Printf("sum=", sum)
	//fmt.Println(len(pipe))
}
