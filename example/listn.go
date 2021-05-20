package main

import "fmt"

// List
func List(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}
