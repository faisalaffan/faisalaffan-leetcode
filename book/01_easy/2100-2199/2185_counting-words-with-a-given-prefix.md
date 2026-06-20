# 2185 — Counting Words With A Given Prefix

## Deskripsi

**Soal:** [2185. Counting Words With A Given Prefix](https://leetcode.com/problems/counting-words-with-a-given-prefix/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2185: Counting Words With a Given Prefix
// https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"pay", "attention", "practice", "attend"}, "at")) // 2
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"leetcode", "win", "loops", "success"}, "code"))   // 0
}

// Time: O(n * m), Space: O(1)
func CountingWordsWithAGivenPrefix(words []string, pref string) int {
	count := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			count++
		}
	}
	return count
}
```
