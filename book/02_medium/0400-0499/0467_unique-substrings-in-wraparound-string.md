# 0467 — Unique Substrings In Wraparound String

## Deskripsi

**Soal:** [0467. Unique Substrings In Wraparound String](https://leetcode.com/problems/unique-substrings-in-wraparound-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (26 letters)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #467: Unique Substrings in Wraparound String
// https://leetcode.com/problems/unique-substrings-in-wraparound-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (26 letters)

import "fmt"

func main() {
	fmt.Println(UniqueSubstringsInWraparoundString("a"))
	fmt.Println(UniqueSubstringsInWraparoundString("cac"))
	fmt.Println(UniqueSubstringsInWraparoundString("zab"))
}

func UniqueSubstringsInWraparoundString(s string) int {
  // Membuat slice untuk menyimpan hasil
	maxLen := make([]int, 26)
	curLen := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if i > 0 && (s[i]-s[i-1] == 1 || (s[i-1] == 'z' && s[i] == 'a')) {
			curLen++
		} else {
			curLen = 1
		}
		idx := s[i] - 'a'
		if curLen > maxLen[idx] {
			maxLen[idx] = curLen
		}
	}

	total := 0
	for _, v := range maxLen {
		total += v
	}
	return total
}
```
