package main

// LeetCode #2870: Minimum Number of Operations to Make Array Empty
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumNumberOfOperationsToMakeArrayEmpty(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	ops := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// Use as many 3s as possible
		ops += c / 3
		if c%3 != 0 {
			ops++
		}
	}

	return ops
}

func main() {
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{2, 3, 3, 3, 3, 2}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 1, 1, 1}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 2, 3}))
}
