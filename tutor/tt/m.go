package main

import (
	"fmt"
	"os"
	"strings"
)
import "time"

// 这是一个我们将要在 Go 协程中运行的函数。`done` 通道
// 将被用于通知其他 Go 协程这个函数已经工作完毕。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦。
	done <- true
}

func main() {
	var a = new(int)
	*a = 100
	fmt.Println(*a)

	var b = make(map[string]int)
	b["测试"] = 100
	fmt.Println(b)
	// 运行一个 worker Go协程，并给予用于通知的通道。
	for key, e := range os.Environ() {
		fmt.Println(key)
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
