# 0720 — Longest Word In Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestWord(words []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L + n log n)  
**Kompleksitas Ruang:** O(n * L)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #720: Longest Word in Dictionary
// https://leetcode.com/problems/longest-word-in-dictionary/
// Difficulty: Medium
// Time: O(n * L + n log n)
// Space: O(n * L)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}

func longestWord(words []string) string {
  // Membuat map (HashMap) — pencarian O(1)
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}

	sort.Strings(words)

	result := ""
	for _, w := range words {
		if len(w) <= len(result) {
			continue
		}
		valid := true
		for i := 1; i < len(w); i++ {
			if !wordSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			result = w
		}
	}

	return result
}
```
