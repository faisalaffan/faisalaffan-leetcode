package main

// LeetCode #1963: Minimum Number of Swaps to Make the String Balanced
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwapsBalanced("][]["))
	fmt.Println(MinSwapsBalanced("]]][[["))
	fmt.Println(MinSwapsBalanced("[]"))
}

// Time: O(n), Space: O(1)
func MinSwapsBalanced(s string) int {
	unmatched := 0
	maxUnmatched := 0
	for _, c := range s {
		if c == '[' {
			unmatched--
		} else {
			unmatched++
		}
		if unmatched > maxUnmatched {
			maxUnmatched = unmatched
		}
	}
	return (maxUnmatched + 1) / 2
}
