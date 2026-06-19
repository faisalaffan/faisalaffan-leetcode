package main

// LeetCode #2974: Minimum Number Game
// https://leetcode.com/problems/minimum-number-game/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: numberGame
	fmt.Println(MinimumNumberGame([]int{5, 4, 2, 3})) // [3, 2, 5, 4]
	fmt.Println(MinimumNumberGame([]int{2, 5}))       // [5, 2]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: numberGame
func MinimumNumberGame(nums []int) []int {
	sort.Ints(nums)
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i += 2 {
		result[i] = nums[i+1]
		result[i+1] = nums[i]
	}
	return result
}
