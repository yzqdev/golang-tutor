package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]

		_, ok := m[another]
		fmt.Println(ok, m, another)
		if ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6}, 6))
}
