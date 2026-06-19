package main

// LeetCode #2531: Make Number of Distinct Characters Equal
// https://leetcode.com/problems/make-number-of-distinct-characters-equal/
// Difficulty: Medium
// Time: O(26^2) | Space: O(26)
// Try swapping one char from word1 with one char from word2.
// Check if distinct counts become equal.

import "fmt"

func main() {
	fmt.Println(isItPossible("ac", "b"))    // false
	fmt.Println(isItPossible("abcc", "aab")) // true
}

func isItPossible(word1 string, word2 string) bool {
	c1, c2 := make([]int, 26), make([]int, 26)
	for _, ch := range word1 {
		c1[ch-'a']++
	}
	for _, ch := range word2 {
		c2[ch-'a']++
	}

	for i := 0; i < 26; i++ {
		if c1[i] == 0 {
			continue
		}
		for j := 0; j < 26; j++ {
			if c2[j] == 0 {
				continue
			}
			// swap char i (from word1) with char j (from word2)
			// decrement c1[i], c2[j]; increment c1[j], c2[i]
			c1[i]--
			c2[j]--
			c1[j]++
			c2[i]++

			d1, d2 := 0, 0
			for k := 0; k < 26; k++ {
				if c1[k] > 0 {
					d1++
				}
				if c2[k] > 0 {
					d2++
				}
			}

			if d1 == d2 {
				return true
			}

			// revert
			c1[i]++
			c2[j]++
			c1[j]--
			c2[i]--
		}
	}
	return false
}
