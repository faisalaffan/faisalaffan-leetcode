# 0648 — Replace Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func replaceWords(dictionary []string, sentence string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * L) where n is number of words in sentence, L is avg word length  |  **Ruang:** O(d) where d is total characters in dictionary

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #648: Replace Words
// https://leetcode.com/problems/replace-words/
// Difficulty: Medium
// Time: O(n * L) where n is number of words in sentence, L is avg word length
// Space: O(d) where d is total characters in dictionary

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}

func replaceWords(dictionary []string, sentence string) string {
  // HashMap: O(1) lookup
	rootSet := make(map[string]bool)
	for _, root := range dictionary {
		rootSet[root] = true
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if rootSet[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
```
