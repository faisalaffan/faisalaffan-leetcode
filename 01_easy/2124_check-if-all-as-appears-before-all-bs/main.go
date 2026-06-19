package main

// LeetCode #2124: Check if All A's Appears Before All B's
// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("aaabbb")) // true
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("abab"))   // false
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("bbb"))    // true
}

// Time: O(n), Space: O(1)
func CheckIfAllAsAppearsBeforeAllBs(s string) bool {
	foundB := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			foundB = true
		} else if s[i] == 'a' && foundB {
			return false
		}
	}
	return true
}
