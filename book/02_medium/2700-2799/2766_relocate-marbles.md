# 2766 — Relocate Marbles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2766: Relocate Marbles
// https://leetcode.com/problems/relocate-marbles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
  // HashMap: O(1) lookup
	positions := make(map[int]bool)
	for _, n := range nums {
		positions[n] = true
	}

  // Linear scan O(n)
	for i := 0; i < len(moveFrom); i++ {
		delete(positions, moveFrom[i])
		positions[moveTo[i]] = true
	}

  // Alokasi slice
	result := make([]int, 0, len(positions))
	for p := range positions {
		result = append(result, p)
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(RelocateMarbles([]int{1, 2, 3}, []int{1}, []int{4}))
	fmt.Println(RelocateMarbles([]int{1, 1, 2, 2}, []int{1, 2}, []int{3, 4}))
}
```
