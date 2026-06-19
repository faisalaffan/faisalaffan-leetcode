package main

// LeetCode #395: Longest Substring with At Least K Repeating Characters
// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
// Difficulty: Medium
// Time: O(n^2) worst case, O(n) average | Space: O(n)

import "fmt"

func longestSubstring(s string, k int) int {
	return longestSubstringHelper(s, 0, len(s), k)
}

func longestSubstringHelper(s string, start, end, k int) int {
	if end-start < k {
		return 0
	}

	// Count frequencies
	freq := [26]int{}
	for i := start; i < end; i++ {
		freq[s[i]-'a']++
	}

	// Find split point where char freq < k
	for i := start; i < end; i++ {
		if freq[s[i]-'a'] < k {
			left := longestSubstringHelper(s, start, i, k)
			right := longestSubstringHelper(s, i+1, end, k)
			if left > right {
				return left
			}
			return right
		}
	}

	return end - start
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestSubstring("aaabb", 3))
	// Expected: 3 ("aaa")

	// Test case 2
	fmt.Println("Test 2:", longestSubstring("ababbc", 2))
	// Expected: 5 ("ababb")

	// Test case 3
	fmt.Println("Test 3:", longestSubstring("aaabbb", 3))
	// Expected: 6
}
