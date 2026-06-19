package main

// LeetCode #2414: Length of the Longest Alphabetical Continuous Substring
// https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Scan, count consecutive chars where s[i] == s[i-1] + 1.

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abacaba")) // 2 ("ab")
	fmt.Println(longestContinuousSubstring("abcde"))   // 5
}

func longestContinuousSubstring(s string) int {
	ans, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
