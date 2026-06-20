# 1859 — Sorting The Sentence

## Deskripsi

**Soal:** [1859. Sorting The Sentence](https://leetcode.com/problems/sorting-the-sentence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func SortSentence(s string) string`

## Solusi Go

```go
package main

// LeetCode #1859: Sorting the Sentence
// https://leetcode.com/problems/sorting-the-sentence/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func SortSentence(s string) string {
	words := strings.Split(s, " ")
  // Membuat slice untuk menyimpan hasil
	result := make([]string, len(words))
	for _, w := range words {
		pos := int(w[len(w)-1] - '0') - 1
		result[pos] = w[:len(w)-1]
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(SortSentence("is2 sentence4 This1 a3"))
	fmt.Println(SortSentence("Myself2 Me1 I4 and3"))
}
```
