# 3090 — Maximum Length Substring With Two Occurrences

## Deskripsi

**Soal:** [3090. Maximum Length Substring With Two Occurrences](https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3090: Maximum Length Substring With Two Occurrences
// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumLengthSubstring
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("bcbbbcba")) // 4
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("aaaa"))      // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maximumLengthSubstring
func MaximumLengthSubstringWithTwoOccurrences(s string) int {
	left := 0
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[byte]int)
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq[s[right]] > 2 {
			freq[s[left]]--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
