package main

// LeetCode #1684: Count the Number of Consistent Strings
// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n + m*k), Space: O(1)
func CountConsistentStrings(allowed string, words []string) int {
	allowedSet := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		allowedSet[allowed[i]] = true
	}
	count := 0
	for _, word := range words {
		consistent := true
		for i := 0; i < len(word); i++ {
			if !allowedSet[word[i]] {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(CountConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(CountConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
