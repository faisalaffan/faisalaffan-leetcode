package main

// LeetCode #3091: Apply Operations to Make Sum of Array Greater Than or Equal to k
// https://leetcode.com/problems/apply-operations-to-make-sum-of-array-greater-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(sqrt(k)) | Space: O(1)

import "fmt"

func minOperations(k int) int {
	ans := k - 1
	for a := 1; a <= k; a++ {
		b := (k + a - 1) / a
		ops := (a - 1) + (b - 1)
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations(11))  // Expected: 5
	fmt.Println(minOperations(1))   // Expected: 0
	fmt.Println(minOperations(5))   // Expected: 3
}
