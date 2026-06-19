package main

// LeetCode #1456: Maximum Number of Vowels in a Substring of Given Length
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxVowels("abciiidef", 3)) // 3

	// Test case 2
	fmt.Println(maxVowels("aeiou", 2)) // 2

	// Test case 3
	fmt.Println(maxVowels("leetcode", 3)) // 2

	// Test case 4
	fmt.Println(maxVowels("rhythms", 4)) // 0
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

// Time: O(n) where n = length of string
// Space: O(1)
func maxVowels(s string, k int) int {
	// Count vowels in first window
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	maxCount := count

	// Sliding window
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			count--
		}
		if isVowel(s[i]) {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
