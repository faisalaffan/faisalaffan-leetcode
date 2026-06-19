package main

// LeetCode #1400: Construct K Palindrome Strings
// https://leetcode.com/problems/construct-k-palindrome-strings/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(canConstruct("annabelle", 2)) // true

	// Test case 2
	fmt.Println(canConstruct("leetcode", 3)) // false

	// Test case 3
	fmt.Println(canConstruct("true", 4)) // true

	// Test case 4
	fmt.Println(canConstruct("yzyzyzyzyzyzyzy", 2)) // true
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 26
func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	// Count characters with odd frequency
	oddCount := 0
	for _, f := range freq {
		if f%2 == 1 {
			oddCount++
		}
	}

	// Each palindrome can have at most 1 odd-count character
	return oddCount <= k
}
