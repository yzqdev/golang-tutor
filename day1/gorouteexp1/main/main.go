package main

import (
	"fmt"
	"golangtutorial/src/go_dev/day1/gorouteexp1/gorout"
)

func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go gorout.Add(100, 300, pipe)
	sum := <-pipe
	fmt.Println(sum)
}
