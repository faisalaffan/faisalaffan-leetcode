# 0884 — Uncommon Words From Two Sentences

## Deskripsi

**Soal:** [0884. Uncommon Words From Two Sentences](https://leetcode.com/problems/uncommon-words-from-two-sentences/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m). Space: O(n + m).  
**Kompleksitas Ruang:** O(n + m).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #884: Uncommon Words from Two Sentences
// https://leetcode.com/problems/uncommon-words-from-two-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour")) // [sweet sour]
	fmt.Println(uncommonFromSentences("apple apple", "banana"))                     // [banana]
}

// uncommonFromSentences returns all uncommon words across two sentences.
// Time: O(n + m). Space: O(n + m).
func uncommonFromSentences(s1 string, s2 string) []string {
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[string]int)
	for _, w := range strings.Fields(s1) {
		count[w]++
	}
	for _, w := range strings.Fields(s2) {
		count[w]++
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	for w, c := range count {
		if c == 1 {
			result = append(result, w)
		}
	}
	return result
}
```
