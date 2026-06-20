# 0819 — Most Common Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func mostCommonWord(paragraph string, banned []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m). Space: O(n + m).  
**Kompleksitas Ruang:** O(n + m).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #819: Most Common Word
// https://leetcode.com/problems/most-common-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})) // "ball"
	fmt.Println(mostCommonWord("a.", []string{}))                                                          // "a"
}

// mostCommonWord finds the most frequent word in paragraph that is not banned.
// Time: O(n + m). Space: O(n + m).
func mostCommonWord(paragraph string, banned []string) string {
  // Membuat map (HashMap) — pencarian O(1)
	bannedSet := make(map[string]bool)
	for _, w := range banned {
		bannedSet[w] = true
	}

	// Normalize: lowercase and split by non-letter characters
	normalized := strings.ToLower(paragraph)
	cleaned := strings.NewReplacer("!", " ", "?", " ", "'", " ", ",", " ", ";", " ", ".", " ").Replace(normalized)

	words := strings.Fields(cleaned)
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[string]int)
	maxCount := 0
	result := ""
	for _, w := range words {
		if bannedSet[w] {
			continue
		}
		counts[w]++
		if counts[w] > maxCount {
			maxCount = counts[w]
			result = w
		}
	}
	return result
}
```
