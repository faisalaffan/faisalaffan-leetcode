package main

// LeetCode #3520: Minimum Threshold for Inversion Pairs Count
// https://leetcode.com/problems/minimum-threshold-for-inversion-pairs-count/
// Difficulty: Medium [Paid]
// Complexity: O(n log n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumThresholdForInversionPairsCount([]int{1, 3, 2, 4}, 1))
	// Test case 2
	fmt.Println("Test 2:", MinimumThresholdForInversionPairsCount([]int{4, 3, 2, 1}, 3))
	// Test case 3
	fmt.Println("Test 3:", MinimumThresholdForInversionPairsCount([]int{1, 2, 3}, 0))
}

func MinimumThresholdForInversionPairsCount(arr []int, threshold int) int {
	// Count inversion pairs (i < j, arr[i] > arr[j])
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	if count >= threshold {
		return 1
	}
	return 0
}
