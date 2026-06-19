package main

// LeetCode #2606: Find the Substring With Maximum Cost
// https://leetcode.com/problems/find-the-substring-with-maximum-cost/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumCostSubstring(s string, chars string, vals []int) int {
	// Build cost map
	cost := make([]int, 26)
	for i := 0; i < 26; i++ {
		cost[i] = i + 1 // 'a' = 1, 'b' = 2, ...
	}
	for i, c := range chars {
		cost[c-'a'] = vals[i]
	}

	// Kadane's algorithm
	maxEnd := 0
	maxSoFar := 0
	for _, ch := range s {
		val := cost[ch-'a']
		maxEnd = maxEnd + val
		if maxEnd < 0 {
			maxEnd = 0
		}
		if maxEnd > maxSoFar {
			maxSoFar = maxEnd
		}
	}
	return maxSoFar
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumCostSubstring("adaa", "d", []int{-1000}))
	// Expected: 2 (choose "aa" for cost 2)

	// Test case 2
	fmt.Println("Test 2:", maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
	// Expected: 0 (empty substring)

	// Test case 3
	fmt.Println("Test 3:", maximumCostSubstring("aabc", "ab", []int{5, 5}))
	// Expected: 10 (choose "aa" with custom values)
}
