package main

// LeetCode #3746: Minimum String Length After Balanced Removals
// https://leetcode.com/problems/minimum-string-length-after-balanced-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumStringLengthAfterBalancedRemovals(s string) int {
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
		} else {
			b++
		}
	}
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aabbab"))
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aaaa"))
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aaabb"))
}
