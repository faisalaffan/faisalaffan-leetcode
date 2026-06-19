package main

// LeetCode #3185: Count Pairs That Form a Complete Day II
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(24)

import "fmt"

func countCompleteDayPairs(hours []int) int64 {
	count := make([]int, 24)
	var ans int64

	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		ans += int64(count[need])
		count[r]++
	}
	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24})) // Expected: 2
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))       // Expected: 3
}
