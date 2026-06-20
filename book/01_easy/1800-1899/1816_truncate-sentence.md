# 1816 — Truncate Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func TruncateSentence(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1816: Truncate Sentence
// https://leetcode.com/problems/truncate-sentence/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func TruncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	return strings.Join(words[:k], " ")
}

func main() {
	fmt.Println(TruncateSentence("Hello how are you Contestant", 4))
	fmt.Println(TruncateSentence("What is the solution to this problem", 4))
	fmt.Println(TruncateSentence("chopper is not a tanuki", 5))
}
```
