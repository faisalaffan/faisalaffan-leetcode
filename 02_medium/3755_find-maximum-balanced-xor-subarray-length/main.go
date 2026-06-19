package main

// LeetCode #3755: Find Maximum Balanced XOR Subarray Length
// https://leetcode.com/problems/find-maximum-balanced-xor-subarray-length/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findMaximumBalancedXorSubarrayLength(nums []int) int {
	type state struct {
		xor  int
		diff int
	}
	first := make(map[state]int)
	// Initial state before first element
	first[state{xor: 0, diff: 0}] = -1

	prefixXor := 0
	diff := 0
	maxLen := 0

	for i, v := range nums {
		prefixXor ^= v
		if v%2 == 1 {
			diff++
		} else {
			diff--
		}

		key := state{xor: prefixXor, diff: diff}
		if pos, ok := first[key]; ok {
			if i-pos > maxLen {
				maxLen = i - pos
			}
		} else {
			first[key] = i
		}
	}

	return maxLen
}

func main() {
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{3, 1, 3, 2, 0}))
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{3, 2, 8, 5, 4, 14, 9, 15}))
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{1, 2, 3, 4, 5}))
}
