package main

// LeetCode #3121: Count the Number of Special Characters II
// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func numberOfSpecialChars(word string) int {
	firstLower := make([]int, 26)
	lastUpper := make([]int, 26)
	for i := range firstLower {
		firstLower[i] = -1
		lastUpper[i] = -1
	}

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'
			if firstLower[idx] == -1 {
				firstLower[idx] = i
			}
		} else {
			idx := ch - 'A'
			lastUpper[idx] = i
		}
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if firstLower[i] != -1 && lastUpper[i] != -1 && firstLower[i] > lastUpper[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))  // Expected: 3
	fmt.Println(numberOfSpecialChars("abc"))       // Expected: 0
	fmt.Println(numberOfSpecialChars("AbBCab"))    // Expected: 0
}
