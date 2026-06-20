# 1268 — Search Suggestions System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func suggestedProducts(products []string, searchWord string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** O(n log n + m * n) where m = len(searchWord)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1268: Search Suggestions System
// https://leetcode.com/problems/search-suggestions-system/
// Difficulty: Medium

// For each prefix of searchWord, return top 3 lexicographically
// smallest products that match the prefix.

// Time: O(n log n + m * n) where m = len(searchWord)
// Space: O(n)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

  // Matriks 2D
	result := make([][]string, len(searchWord))

	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		suggestions := make([]string, 0)

		for _, p := range products {
			if len(p) >= i && p[:i] == prefix {
				suggestions = append(suggestions, p)
				if len(suggestions) == 3 {
					break
				}
			}
		}

		result[i-1] = suggestions
	}

	return result
}

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"
	result := suggestedProducts(products, searchWord)
	for i, r := range result {
		fmt.Printf("prefix %q: %v\n", searchWord[:i+1], r)
	}
}
```
