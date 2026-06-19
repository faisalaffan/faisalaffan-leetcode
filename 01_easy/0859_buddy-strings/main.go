package main

// LeetCode #859: Buddy Strings
// https://leetcode.com/problems/buddy-strings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))   // true
	fmt.Println(buddyStrings("ab", "ab"))   // false
	fmt.Println(buddyStrings("aa", "aa"))   // true
	fmt.Println(buddyStrings("abcd", "badc")) // false
}

// buddyStrings checks if swapping two letters in s makes it equal to goal.
// Time: O(n). Space: O(1).
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		// Need at least one duplicate character to swap
		seen := make(map[byte]bool)
		for i := range s {
			if seen[s[i]] {
				return true
			}
			seen[s[i]] = true
		}
		return false
	}
	diff := make([]int, 0)
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}
