package main

// LeetCode #3365: Rearrange K Substrings to Form Target String
// https://leetcode.com/problems/rearrange-k-substrings-to-form-target-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossibleToRearrange("abcd", "cdab", 2)) // true
	fmt.Println(isPossibleToRearrange("aabb", "bbaa", 2)) // true
	fmt.Println(isPossibleToRearrange("abcd", "acbd", 2)) // false
}

func isPossibleToRearrange(s string, t string, k int) bool {
	n := len(s)
	m := n / k

	cnt := make(map[string]int)
	for i := 0; i < n; i += m {
		cnt[s[i:i+m]]++
		cnt[t[i:i+m]]--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
