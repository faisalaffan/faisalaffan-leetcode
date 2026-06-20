# 1048 — Longest String Chain

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestStrChain(words []string) int
```

> **💡 Hint:** Sort by length, DP with hash map.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * L^2) where L is max word length  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1048: Longest String Chain
// https://leetcode.com/problems/longest-string-chain/
// Difficulty: Medium
//
// Approach: Sort by length, DP with hash map.
//           For each word, check all possible predecessors by removing one char.
// Time: O(n * L^2) where L is max word length
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"})) // 4
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"})) // 5
}

func longestStrChain(words []string) int {
  // Custom sort dengan comparator
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

  // Membuat map (HashMap) — pencarian O(1)
	dp := make(map[string]int)
	result := 1

	for _, w := range words {
		best := 1
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(w); i++ {
			pred := w[:i] + w[i+1:]
			if val, ok := dp[pred]; ok {
				if val+1 > best {
					best = val + 1
				}
			}
		}
		dp[w] = best
		if best > result {
			result = best
		}
	}

	return result
}
```
