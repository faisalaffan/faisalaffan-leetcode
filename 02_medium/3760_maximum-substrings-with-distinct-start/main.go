package main

// LeetCode #3760: Maximum Substrings With Distinct Start
// https://leetcode.com/problems/maximum-substrings-with-distinct-start/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSubstringsWithDistinctStart(s string) int {
	seen := [26]bool{}
	ans := 0
	for _, ch := range s {
		idx := ch - 'a'
		if !seen[idx] {
			seen[idx] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumSubstringsWithDistinctStart("abacaba"))
	fmt.Println(maximumSubstringsWithDistinctStart("aaaa"))
	fmt.Println(maximumSubstringsWithDistinctStart("abc"))
}
