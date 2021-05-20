package main

import (
	"fmt"
)

func main() {
	totalCount := 0

	for n := 2; n < 1000; n++ {
		m := n
		for i := 1; i < n; i++ {
			if n%i == 0 {
				m -= i
			}
		}
	}
	fmt.Println("嗨客网(www.haicoder.net)")
	fmt.Println("共", totalCount, "种方案")
}
