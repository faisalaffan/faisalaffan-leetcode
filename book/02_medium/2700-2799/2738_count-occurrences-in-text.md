# 2738 — Count Occurrences In Text

## Deskripsi

**Soal:** [2738. Count Occurrences In Text](https://leetcode.com/problems/count-occurrences-in-text/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CountOccurrencesInText(text string, word string) int`

## Solusi Go

```go
package main

// LeetCode #2738: Count Occurrences in Text
// https://leetcode.com/problems/count-occurrences-in-text/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func CountOccurrencesInText(text string, word string) int {
	count := 0
	words := strings.Fields(text)
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountOccurrencesInText("hello world hello", "hello"))
	fmt.Println(CountOccurrencesInText("this is a test test this", "test"))
	fmt.Println(CountOccurrencesInText("unique", "none"))
}
```
