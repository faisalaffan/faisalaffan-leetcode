package main

// LeetCode #926: Flip String to Monotone Increasing
// https://leetcode.com/problems/flip-string-to-monotone-increasing/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func minFlipsMonoIncr(s string) int {
	ones, flips := 0, 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		} else {
			flips = min(flips+1, ones)
		}
	}
	return flips
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}
