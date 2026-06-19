package main

// LeetCode #1806: Minimum Number of Operations to Reinitialize a Permutation
// https://leetcode.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func reinitializePermutation(n int) int {
	ops := 0
	i := 1

	for {
		ops++
		if i%2 == 0 {
			i /= 2
		} else {
			i = n/2 + (i-1)/2
		}
		if i == 1 {
			break
		}
	}
	return ops
}

func main() {
	fmt.Println(reinitializePermutation(2))  // Expected: 1
	fmt.Println(reinitializePermutation(4))  // Expected: 2
	fmt.Println(reinitializePermutation(6))  // Expected: 4
}
