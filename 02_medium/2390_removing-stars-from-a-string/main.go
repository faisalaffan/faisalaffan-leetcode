package main

// LeetCode #2390: Removing Stars From a String
// https://leetcode.com/problems/removing-stars-from-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Stack: push non-star, pop on star.

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e")) // "lecoe"
	fmt.Println(removeStars("erase*****"))  // ""
}

func removeStars(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
