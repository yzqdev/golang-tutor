package main

import "fmt"

func TwoSum(nums []int, target int) []int {
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
