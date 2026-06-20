# 2782 — Number Of Unique Categories

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumberOfUniqueCategories(categories []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2782: Number of Unique Categories
// https://leetcode.com/problems/number-of-unique-categories/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func NumberOfUniqueCategories(categories []string) int {
  // HashMap: O(1) lookup
	seen := make(map[string]bool)
	for _, c := range categories {
		seen[c] = true
	}
	return len(seen)
}

func main() {
	fmt.Println(NumberOfUniqueCategories([]string{"a", "b", "a", "c"}))
	fmt.Println(NumberOfUniqueCategories([]string{"x", "x", "x"}))
	fmt.Println(NumberOfUniqueCategories([]string{}))
}
```
