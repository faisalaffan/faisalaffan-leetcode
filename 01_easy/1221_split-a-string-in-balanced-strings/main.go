package main

// LeetCode #1221: Split a String in Balanced Strings
// https://leetcode.com/problems/split-a-string-in-balanced-strings/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // 4
	fmt.Println(balancedStringSplit("RLLLLRRRLR")) // 3
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // 1
}

// LeetCode submission: balancedStringSplit
func balancedStringSplit(s string) int {
	count, ans := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans++
		}
	}
	return ans
}
