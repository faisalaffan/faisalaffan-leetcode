# 2114 — Maximum Number Of Words Found In Sentences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumNumberOfWordsFoundInSentences(sentences []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2114: Maximum Number of Words Found in Sentences
// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})) // 6
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"please wait", "continue to fight", "continue to win"}))                              // 3
}

// Time: O(n * m), Space: O(1)
func MaximumNumberOfWordsFoundInSentences(sentences []string) int {
	maxWords := 0
	for _, s := range sentences {
		count := strings.Count(s, " ") + 1
		if count > maxWords {
			maxWords = count
		}
	}
	return maxWords
}
```
