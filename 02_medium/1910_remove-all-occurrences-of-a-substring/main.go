package main

// LeetCode #1910: Remove All Occurrences of a Substring
// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(RemoveOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(RemoveOccurrences("axxxxyyyyb", "xy"))
	fmt.Println(RemoveOccurrences("aabababa", "aba"))
}

// Time: O(n*m) where n = len(s), m = len(part), Space: O(n)
func RemoveOccurrences(s string, part string) string {
	stack := make([]byte, 0)
	m := len(part)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= m && string(stack[len(stack)-m:]) == part {
			stack = stack[:len(stack)-m]
		}
	}
	return string(stack)
}
