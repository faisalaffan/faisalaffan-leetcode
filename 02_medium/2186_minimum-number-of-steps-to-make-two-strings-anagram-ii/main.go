package main

// LeetCode #2186: Minimum Number of Steps to Make Two Strings Anagram II
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)

import "fmt"

func minSteps(s string, t string) int {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}
	for _, ch := range t {
		count[ch-'a']--
	}

	steps := 0
	for _, c := range count {
		if c > 0 {
			steps += c
		} else {
			steps -= c
		}
	}
	return steps
}

func main() {
	// Test case 1
	fmt.Println(minSteps("leetcode", "coats"))
	// Expected: 7

	// Test case 2
	fmt.Println(minSteps("night", "thing"))
	// Expected: 0

	// Test case 3
	fmt.Println(minSteps("aba", "bab"))
	// Expected: 2
}
