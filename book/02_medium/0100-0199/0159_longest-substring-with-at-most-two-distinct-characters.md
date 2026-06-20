# 0159 — Longest Substring With At Most Two Distinct Characters

## Deskripsi

**Soal:** [0159. Longest Substring With At Most Two Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func lengthOfLongestSubstringTwoDistinct(s string) int`

## Solusi Go

```go
package main

// LeetCode #159: Longest Substring with At Most Two Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
  // Membuat map untuk pencarian O(1): key → value
	charCount := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		for len(charCount) > 2 {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
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
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abc"))
}
```
