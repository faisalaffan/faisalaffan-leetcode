package main

// LeetCode #2598: Smallest Missing Non-negative Integer After Operations
// https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findSmallestInteger(nums []int, value int) int {
	freq := make([]int, value)
	for _, v := range nums {
		// Map to [0, value-1] range
		rem := ((v % value) + value) % value
		freq[rem]++
	}

	for i := 0; ; i++ {
		if freq[i%value] == 0 {
			return i
		}
		freq[i%value]--
	}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", findSmallestInteger([]int{1, 2, 3, 4, 5}, 2))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", findSmallestInteger([]int{0, 0, 0, 0}, 1))
	// Expected: 4
}
