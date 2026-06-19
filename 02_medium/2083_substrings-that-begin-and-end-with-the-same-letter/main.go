package main

// LeetCode #2083: Substrings That Begin and End With the Same Letter
// https://leetcode.com/problems/substrings-that-begin-and-end-with-the-same-letter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfSubstrings(s string) int64 {
	freq := make([]int64, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	var result int64 = 0
	for _, f := range freq {
		result += f * (f + 1) / 2
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubstrings("abc"))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfSubstrings("abacaba"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", numberOfSubstrings("aa"))
	// Expected: 3
}
