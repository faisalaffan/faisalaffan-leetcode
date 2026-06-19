package main

// LeetCode #2560: House Robber IV
// https://leetcode.com/problems/house-robber-iv/
// Difficulty: Medium
// Time: O(n log max) | Space: O(1)

import "fmt"

func minCapability(nums []int, k int) int {
	canRob := func(cap int) bool {
		count := 0
		i := 0
		for i < len(nums) {
			if nums[i] <= cap {
				count++
				i += 2 // Skip adjacent house
			} else {
				i++
			}
		}
		return count >= k
	}

	left, right := nums[0], nums[0]
	for _, v := range nums {
		if v < left {
			left = v
		}
		if v > right {
			right = v
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if canRob(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCapability([]int{2, 3, 5, 9}, 2))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", minCapability([]int{2, 7, 9, 3, 1}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minCapability([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	// Expected: 5
}
