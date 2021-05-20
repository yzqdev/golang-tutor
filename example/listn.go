package main

import "fmt"

func main() {
	list(10)

}

/**
 *求出所有和为n的两个数
 */
func list(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}
