package main

// LeetCode #3084: Count Substrings Starting and Ending with Given Character
// https://leetcode.com/problems/count-substrings-starting-and-ending-with-given-character/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubstringsStartingEnding("abada", 'a'))
	fmt.Println(countSubstringsStartingEnding("zzz", 'z'))
}

func countSubstringsStartingEnding(s string, c byte) int64 {
	cnt := int64(0)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt++
		}
	}
	return cnt * (cnt + 1) / 2
}
