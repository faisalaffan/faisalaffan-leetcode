package main

// LeetCode #3113: Find the Number of Subarrays Where Boundary Elements Are Maximum
// https://leetcode.com/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/
// Difficulty: Hard
// Time: O(n) | Space: O(n)

import "fmt"

func numberOfSubarrays(nums []int) int64 {
	n := len(nums)
	stack := make([]int, 0)               // monotonic decreasing stack (indices)
	count := make([]int64, n)             // count of same-value elements at each position
	var result int64 = 0

	for i := 0; i < n; i++ {
		// Pop elements smaller than nums[i] (they can't be left boundaries)
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}

		// If top of stack has the same value, chain the count
		if len(stack) > 0 && nums[stack[len(stack)-1]] == nums[i] {
			count[i] = count[stack[len(stack)-1]] + 1
		} else {
			count[i] = 1
		}

		stack = append(stack, i)
		result += count[i]
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubarrays([]int{1, 4, 3, 3, 2}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", numberOfSubarrays([]int{3, 3, 3}))
	// Expected: 6

	// Test case 3
	fmt.Println("Test 3:", numberOfSubarrays([]int{1}))
	// Expected: 1
}
