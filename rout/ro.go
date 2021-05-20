package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	c := make(chan int, 2)
	c <- 1

	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
