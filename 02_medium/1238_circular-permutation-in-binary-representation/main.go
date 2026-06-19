package main

import (
	"fmt"
)

// LeetCode #1238: Circular Permutation in Binary Representation
// https://leetcode.com/problems/circular-permutation-in-binary-representation/
// Difficulty: Medium

// Generate Gray code sequence starting with start.
// Gray code: consecutive numbers differ by exactly 1 bit.

// Time: O(2^n)
// Space: O(2^n)

func circularPermutation(n int, start int) []int {
	size := 1 << n
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = start ^ i ^ (i >> 1)
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [3 2 0 1] or similar)\n", circularPermutation(2, 3))
	fmt.Printf("%v\n", circularPermutation(3, 2))
}
