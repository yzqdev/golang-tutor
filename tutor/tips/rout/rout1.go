package rout

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func GetWg() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func Add(a int, b int, c chan int) {
	sum := a + b
	c <- sum
}
func GetChan() {
	var pipe chan int
	pipe = make(chan int, 1)
	go Add(100, 300, pipe)
	sum := <-pipe
	fmt.Println(sum)
}
