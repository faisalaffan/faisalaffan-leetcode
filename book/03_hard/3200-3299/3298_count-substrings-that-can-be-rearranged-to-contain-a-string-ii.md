# 3298 — Count Substrings That Can Be Rearranged To Contain A String Ii

## Deskripsi

**Soal:** [3298. Count Substrings That Can Be Rearranged To Contain A String Ii](https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

## Solusi Go

```go
package main

// LeetCode #3298: Count Substrings That Can Be Rearranged to Contain a String II
// https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-ii/
// Difficulty: Hard
//
// Count substrings of word1 that contain all the characters of word2
// (with at least the required frequency) after rearrangement.
// Since rearrangement is allowed, a substring is valid iff for every
// character c, freq_in_substring[c] >= freq_in_word2[c].
//
// Sliding window: for each right pointer, find the minimal left pointer
// such that the window satisfies the frequency condition. Then all
// substrings ending at right with start <= left are valid.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(validSubstringCount("abcabc", "abc"))
	// Example 2
	fmt.Println(validSubstringCount("aabbcc", "abc"))
	// Example 3: word2 longer than word1
	fmt.Println(validSubstringCount("abc", "abcd"))
	// Example 4: single char
	fmt.Println(validSubstringCount("aaaa", "a"))
	// Example 5: exact match
	fmt.Println(validSubstringCount("leetcode", "code"))
	// Example 6: many repeated chars
	fmt.Println(validSubstringCount("aaabbbccc", "abc"))
}

func validSubstringCount(word1 string, word2 string) int64 {
	n := len(word1)

	// Frequency of characters needed from word2.
	need := [26]int{}
	for _, ch := range word2 {
		need[ch-'a']++
	}

	// Count how many distinct characters are required.
	required := 0
	for _, f := range need {
		if f > 0 {
			required++
		}
	}

	// Sliding window.
	have := [26]int{}
	formed := 0 // Number of chars meeting the requirement.
	left := 0
	var count int64

	for right := 0; right < n; right++ {
		c := word1[right] - 'a'
		have[c]++
		if have[c] == need[c] {
			formed++
		}

		// While the window [left..right] satisfies the condition,
		// count all substrings ending at right with any start <= left.
		for left <= right && formed == required {
			// All left' in [0, left] give valid windows [left'..right].
			count += int64(left + 1)

			// Contract from the left.
			c2 := word1[left] - 'a'
			have[c2]--
			if have[c2] < need[c2] {
				formed--
			}
			left++
		}
	}

	return count
}
```
