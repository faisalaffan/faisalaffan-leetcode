# 0954 — Array Of Doubled Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func canReorderDoubled(arr []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #954: Array of Doubled Pairs
// https://leetcode.com/problems/array-of-doubled-pairs/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func canReorderDoubled(arr []int) bool {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

  // Alokasi slice
	keys := make([]int, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
  // Custom sort
	sort.Slice(keys, func(i, j int) bool {
		return abs(keys[i]) < abs(keys[j])
	})

	for _, v := range keys {
		if freq[v] > freq[2*v] {
			return false
		}
		freq[2*v] -= freq[v]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{1, 2, 4, 16, 8, 4}))
}
```
