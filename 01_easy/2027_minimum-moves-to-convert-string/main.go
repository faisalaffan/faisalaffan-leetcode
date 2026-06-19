package main

// LeetCode #2027: Minimum Moves to Convert String
// https://leetcode.com/problems/minimum-moves-to-convert-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToConvertString("XXX"))       // 1
	fmt.Println(MinimumMovesToConvertString("XXOX"))      // 2
	fmt.Println(MinimumMovesToConvertString("OOOO"))      // 0
}

// Time: O(n), Space: O(1)
func MinimumMovesToConvertString(s string) int {
	moves := 0
	i := 0
	for i < len(s) {
		if s[i] == 'X' {
			moves++
			i += 3
		} else {
			i++
		}
	}
	return moves
}
