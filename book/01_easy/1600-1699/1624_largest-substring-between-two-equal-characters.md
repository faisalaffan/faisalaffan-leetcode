# 1624 — Largest Substring Between Two Equal Characters

## Deskripsi

**Soal:** [1624. Largest Substring Between Two Equal Characters](https://leetcode.com/problems/largest-substring-between-two-equal-characters/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1) (since only 26 letters)  
**Kompleksitas Ruang:** O(1) (since only 26 letters)

**Algoritma:** —

**Fungsi Solusi:** `func MaxLengthBetweenEqualCharacters(s string) int`

## Solusi Go

```go
package main

// LeetCode #1624: Largest Substring Between Two Equal Characters
// https://leetcode.com/problems/largest-substring-between-two-equal-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1) (since only 26 letters)
func MaxLengthBetweenEqualCharacters(s string) int {
  // Membuat map untuk pencarian O(1): key → value
	firstIndex := make(map[rune]int)
	maxLen := -1
	for i, ch := range s {
		if idx, exists := firstIndex[ch]; exists {
			if i-idx-1 > maxLen {
				maxLen = i - idx - 1
			}
		} else {
			firstIndex[ch] = i
		}
	}
	return maxLen
}

func main() {
	fmt.Println(MaxLengthBetweenEqualCharacters("aa"))
	fmt.Println(MaxLengthBetweenEqualCharacters("abca"))
	fmt.Println(MaxLengthBetweenEqualCharacters("cbzxy"))
}
```
