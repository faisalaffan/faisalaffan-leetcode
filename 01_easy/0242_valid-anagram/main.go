package main

// LeetCode #242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed 26 chars)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	return count == [26]int{}
}

func main() {
	fmt.Println(IsAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram("rat", "car"))
}
