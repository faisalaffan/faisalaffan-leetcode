# 1054 — Distant Barcodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func rearrangeBarcodes(barcodes []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1054: Distant Barcodes
// https://leetcode.com/problems/distant-barcodes/
// Difficulty: Medium
//
// Approach: Count frequencies, place most frequent in even indices, then odd
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 2, 2, 2})) // [1,2,1,2,1,2] or similar
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3})) // valid rearrangement
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, b := range barcodes {
		freq[b]++
	}

	type pair struct {
		val   int
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for val, count := range freq {
		pairs = append(pairs, pair{val, count})
	}
  // Custom sort
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

  // Alokasi slice
	result := make([]int, n)
	idx := 0

	for _, p := range pairs {
		for k := 0; k < p.count; k++ {
			result[idx] = p.val
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return result
}
```
