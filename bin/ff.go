package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Executable()

	fmt.Println(path)
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()
}
