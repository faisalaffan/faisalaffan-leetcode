package main

// LeetCode #205: Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed ASCII chars)
func IsIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	fmt.Println(IsIsomorphic("egg", "add"))
	fmt.Println(IsIsomorphic("foo", "bar"))
	fmt.Println(IsIsomorphic("paper", "title"))
}
