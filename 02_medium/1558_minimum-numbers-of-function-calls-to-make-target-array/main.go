package main

// LeetCode #1558: Minimum Numbers of Function Calls to Make Target Array
// https://leetcode.com/problems/minimum-numbers-of-function-calls-to-make-target-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperations([]int{1, 5}))
	fmt.Println(MinOperations([]int{2, 2}))
	fmt.Println(MinOperations([]int{4, 2, 5}))
}

func MinOperations(nums []int) int {
	// Time: O(N), Space: O(1)
	// Operations: (1) increment one element, (2) double all elements
	// Count total increments (set bits) + maximum number of doublings (highest bit position)
	increments := 0
	maxDoublings := 0

	for _, num := range nums {
		// Count bits (increment operations)
		bits := 0
		pos := 0
		for n := num; n > 0; n >>= 1 {
			if n&1 == 1 {
				bits++
			}
			pos++
		}
		increments += bits
		if pos-1 > maxDoublings {
			maxDoublings = pos - 1
		}
	}

	return increments + maxDoublings
}
