package main

// LeetCode #1844: Replace All Digits with Characters
// https://leetcode.com/problems/replace-all-digits-with-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReplaceDigits(s string) string {
	result := []byte(s)
	for i := 1; i < len(s); i += 2 {
		result[i] = s[i-1] + (s[i] - '0')
	}
	return string(result)
}

func main() {
	fmt.Println(ReplaceDigits("a1c1e1"))
	fmt.Println(ReplaceDigits("a1b2c3d4e"))
}
