# 0524 — Longest Word In Dictionary Through Deleting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindLongestWord(s string, dictionary []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L) where n = len(dictionary), L = max(len(s), len(word))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #524: Longest Word in Dictionary through Deleting
// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/
// Difficulty: Medium
// Time: O(n * L) where n = len(dictionary), L = max(len(s), len(word))
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(FindLongestWord("abpcplea", []string{"a", "b", "c"}))
}

func FindLongestWord(s string, dictionary []string) string {
	// Sort by length desc, then lexicographically
  // Custom sort dengan comparator
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		}
		return dictionary[i] < dictionary[j]
	})

	for _, word := range dictionary {
		if isSubsequence(word, s) {
			return word
		}
	}

	return ""
}

func isSubsequence(word, s string) bool {
	i := 0
	for j := 0; i < len(word) && j < len(s); j++ {
		if word[i] == s[j] {
			i++
		}
	}
	return i == len(word)
}
```
