package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	//go语言没while语句
	fmt.Println(sum)
}
