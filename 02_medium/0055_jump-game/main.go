package main

// LeetCode #55: Jump Game
// https://leetcode.com/problems/jump-game/
// Difficulty: Medium

import "fmt"

func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		if i+nums[i] > reachable {
			reachable = i + nums[i]
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true

	// Test case 2
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false

	// Test case 3
	fmt.Println(canJump([]int{0})) // true
}

// Time: O(n) | Space: O(1)
