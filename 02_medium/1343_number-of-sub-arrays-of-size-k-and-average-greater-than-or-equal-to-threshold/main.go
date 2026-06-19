package main

// LeetCode #1343: Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4)) // 3

	// Test case 2
	fmt.Println(numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5)) // 6

	// Test case 3
	fmt.Println(numOfSubarrays([]int{1, 1, 1, 1, 1}, 1, 0)) // 5
}

// Time: O(n) where n = len(arr)
// Space: O(1)
func numOfSubarrays(arr []int, k int, threshold int) int {
	targetSum := k * threshold
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	count := 0
	if windowSum >= targetSum {
		count++
	}

	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum >= targetSum {
			count++
		}
	}

	return count
}
