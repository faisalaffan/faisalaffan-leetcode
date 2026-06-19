package main

// LeetCode #3223: Minimum Length of String After Operations
// https://leetcode.com/problems/minimum-length-of-string-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func minimumLength(s string) int {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	ans := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 {
			ans += 2
		} else {
			ans += 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumLength("abaacbcbb")) // Expected: 5
	fmt.Println(minimumLength("aa"))         // Expected: 2
}
