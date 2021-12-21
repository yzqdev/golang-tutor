package main

import (
	"fmt"
	"os"
)

func main1() {
	path, _ := os.Executable()

	fmt.Println(path)
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	fmt.Println(file.Chdir())
	// 关闭文件
	defer file.Close()
}
