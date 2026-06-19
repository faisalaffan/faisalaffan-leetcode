package main

// LeetCode #2340: Minimum Adjacent Swaps to Make a Valid Array
// https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumSwaps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	minIdx, maxIdx := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] >= nums[maxIdx] {
			maxIdx = i
		}
	}

	swaps := minIdx + (n - 1 - maxIdx)
	if minIdx > maxIdx {
		swaps--
	}
	return swaps
}

func main() {
	// Test case 1
	fmt.Println(minimumSwaps([]int{3, 4, 5, 5, 3, 1}))
	// Expected: 6

	// Test case 2
	fmt.Println(minimumSwaps([]int{1, 2, 3, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println(minimumSwaps([]int{2, 1}))
	// Expected: 1
}
