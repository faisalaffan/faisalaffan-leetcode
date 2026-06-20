# 0340 — Longest Substring With At Most K Distinct Characters

## Deskripsi

**Soal:** [0340. Longest Substring With At Most K Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

**Algoritma:** —

**Fungsi Solusi:** `func lengthOfLongestSubstringKDistinct(s string, k int) int`

## Solusi Go

```go
package main

// LeetCode #340: Longest Substring with At Most K Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(k)

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for len(freq) > k {
			freq[s[left]]--
			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lengthOfLongestSubstringKDistinct("eceba", 2))
	// Expected: 3 ("ece")

	// Test case 2
	fmt.Println("Test 2:", lengthOfLongestSubstringKDistinct("aa", 1))
	// Expected: 2 ("aa")

	// Test case 3
	fmt.Println("Test 3:", lengthOfLongestSubstringKDistinct("a", 0))
	// Expected: 0
}
```
