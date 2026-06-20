# 1858 — Longest Word With All Prefixes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestWord(words []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n log n + total chars), Space: O(total unique prefixes)  
**Kompleksitas Ruang:** O(total unique prefixes)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1858: Longest Word With All Prefixes
// https://leetcode.com/problems/longest-word-with-all-prefixes/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestWord([]string{"k", "ki", "kir", "kira", "kiran"}))
	fmt.Println(LongestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Println(LongestWord([]string{"abc", "ab", "a"}))
}

// Time: O(n log n + total chars), Space: O(total unique prefixes)
func LongestWord(words []string) string {
  // Membuat map (HashMap) — pencarian O(1)
	prefixSet := make(map[string]bool)
	for _, w := range words {
		prefixSet[w] = true
	}

  // Custom sort dengan comparator
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) > len(words[j])
		}
		return words[i] < words[j]
	})

	for _, w := range words {
		valid := true
		for i := 1; i <= len(w); i++ {
			if !prefixSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			return w
		}
	}
	return ""
}
```
