package main

// LeetCode #2063: Vowels of All Substrings
// https://leetcode.com/problems/vowels-of-all-substrings/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countVowels(word string) int64 {
	n := len(word)
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	var result int64 = 0

	for i := 0; i < n; i++ {
		if vowels[word[i]] {
			// Number of substrings containing word[i]
			// = (i+1) * (n-i)
			result += int64(i+1) * int64(n-i)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countVowels("aba"))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", countVowels("abc"))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", countVowels("no"))
	// Expected: 0
}
