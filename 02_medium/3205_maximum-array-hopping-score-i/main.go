package main

// LeetCode #3205: Maximum Array Hopping Score I
// https://leetcode.com/problems/maximum-array-hopping-score-i/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxScore(nums []int) int {
	ans := 0
	mx := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > mx {
			mx = nums[i]
		}
		ans += mx
	}
	return ans
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5})) // Expected: 14 (2+3+4+5)
	fmt.Println(maxScore([]int{5, 4, 3, 2, 1})) // Expected: 4 (1+1+1+1)
	fmt.Println(maxScore([]int{1, 5, 2, 6, 3})) // Expected: 15 (5+6+6+6)
}
