package main

// LeetCode #387: First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FirstUniqueCharacterInAString(s string) int {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	for i, c := range s {
		if count[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstUniqueCharacterInAString("leetcode"))
	fmt.Println(FirstUniqueCharacterInAString("loveleetcode"))
	fmt.Println(FirstUniqueCharacterInAString("aabb"))
}
