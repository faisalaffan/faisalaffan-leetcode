package main

// LeetCode #438: Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	pFreq := [26]int{}
	for _, ch := range p {
		pFreq[ch-'a']++
	}

	sFreq := [26]int{}
	result := []int{}

	// Initialize first window
	for i := 0; i < len(p); i++ {
		sFreq[s[i]-'a']++
	}
	if sFreq == pFreq {
		result = append(result, 0)
	}

	// Slide window
	for i := len(p); i < len(s); i++ {
		sFreq[s[i]-'a']++
		sFreq[s[i-len(p)]-'a']--
		if sFreq == pFreq {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findAnagrams("cbaebabacd", "abc"))
	// Expected: [0, 6]

	// Test case 2
	fmt.Println("Test 2:", findAnagrams("abab", "ab"))
	// Expected: [0, 1, 2]

	// Test case 3
	fmt.Println("Test 3:", findAnagrams("a", "ab"))
	// Expected: []
}
