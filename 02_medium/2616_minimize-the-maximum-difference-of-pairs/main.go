package main

// LeetCode #2616: Minimize the Maximum Difference of Pairs
// https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/
// Difficulty: Medium
// Time: O(n log n + n log maxDiff) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)

	canForm := func(mid int) bool {
		count := 0
		i := 0
		for i < n-1 {
			if nums[i+1]-nums[i] <= mid {
				count++
				i += 2
			} else {
				i++
			}
			if count >= p {
				return true
			}
		}
		return false
	}

	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2
		if canForm(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minimizeMax([]int{4, 2, 1, 2}, 1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minimizeMax([]int{3, 5, 2, 8, 1, 9}, 3))
	// Expected: 2
}
