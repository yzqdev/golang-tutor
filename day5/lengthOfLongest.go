package main

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}
