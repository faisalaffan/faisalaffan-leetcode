package main

// LeetCode #484: Find Permutation
// https://leetcode.com/problems/find-permutation/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindPermutation("I"))
	fmt.Println(FindPermutation("DI"))
}

func FindPermutation(s string) []int {
	n := len(s) + 1
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}

	// Reverse contiguous segments for each 'D'
	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			j := i
			for j < len(s) && s[j] == 'D' {
				j++
			}
			// Reverse segment from i to j
			left, right := i, j
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			i = j
		}
	}

	return result
}
