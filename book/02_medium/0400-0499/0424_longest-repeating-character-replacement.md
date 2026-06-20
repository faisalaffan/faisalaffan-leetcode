# 0424 — Longest Repeating Character Replacement

## Deskripsi

**Soal:** [0424. Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func characterReplacement(s string, k int) int`

## Solusi Go

```go
package main

// LeetCode #424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left, maxFreq, maxLen := 0, 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]-'A']++
		if freq[s[right]-'A'] > maxFreq {
			maxFreq = freq[s[right]-'A']
		}

		// Window size - maxFreq = chars to replace
		for right-left+1-maxFreq > k {
			freq[s[left]-'A']--
			left++
			// Recompute maxFreq (or keep old - it's safe)
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", characterReplacement("ABAB", 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", characterReplacement("AABABBA", 1))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", characterReplacement("AAAA", 2))
	// Expected: 4
}
```
