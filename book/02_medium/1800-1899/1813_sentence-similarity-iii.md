# 1813 — Sentence Similarity Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func areSentencesSimilar(sentence1 string, sentence2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m), Space: O(n + m)  |  **Ruang:** O(n + m)


## 💻 Solusi Go

```go
package main

// LeetCode #1813: Sentence Similarity III
// https://leetcode.com/problems/sentence-similarity-iii/
// Difficulty: Medium
// Time: O(n + m), Space: O(n + m)

import (
	"fmt"
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	// Make words1 the shorter one
	if len(words1) > len(words2) {
		words1, words2 = words2, words1
	}

	i, j := 0, len(words1)-1
	k, l := 0, len(words2)-1

	// Match from beginning
	for i < len(words1) && k <= l && words1[i] == words2[k] {
		i++
		k++
	}

	// Match from end
	for j >= i && l >= k && words1[j] == words2[l] {
		j--
		l--
	}

	return i > j // All words in shorter sentence matched
}

func main() {
	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley")) // Expected: true
	fmt.Println(areSentencesSimilar("of", "A lot of words")) // Expected: false
	fmt.Println(areSentencesSimilar("Eating right now", "Eating")) // Expected: true
}
```
