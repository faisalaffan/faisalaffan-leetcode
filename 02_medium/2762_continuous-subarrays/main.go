package main

// LeetCode #2762: Continuous Subarrays
// https://leetcode.com/problems/continuous-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ContinuousSubarrays(nums []int) int64 {
	n := len(nums)
	var result int64
	left := 0

	// Track min and max using deques via slice
	// minDeque stores indices with increasing values
	// maxDeque stores indices with decreasing values
	minDeque := make([]int, 0)
	maxDeque := make([]int, 0)

	for right := 0; right < n; right++ {
		// Maintain minDeque
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, right)

		// Maintain maxDeque
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, right)

		// Shrink window if condition violated
		for nums[maxDeque[0]]-nums[minDeque[0]] > 2 {
			if minDeque[0] == left {
				minDeque = minDeque[1:]
			}
			if maxDeque[0] == left {
				maxDeque = maxDeque[1:]
			}
			left++
		}

		result += int64(right - left + 1)
	}

	return result
}

func main() {
	fmt.Println(ContinuousSubarrays([]int{5, 4, 2, 4}))
	fmt.Println(ContinuousSubarrays([]int{1, 2, 3}))
}
