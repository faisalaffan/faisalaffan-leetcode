# 2175 — The Change In Global Rankings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func globalRankings(pointsBefore []int, pointsAfter []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2175: The Change in Global Rankings
// https://leetcode.com/problems/the-change-in-global-rankings/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func globalRankings(pointsBefore []int, pointsAfter []int) []int {
	n := len(pointsBefore)

	// Create sorted list of (points, originalIndex) for before
	type pair struct {
		points int
		idx    int
	}
	before := make([]pair, n)
	for i := 0; i < n; i++ {
		before[i] = pair{pointsBefore[i], i}
	}
  // Custom sort
	sort.Slice(before, func(i, j int) bool {
		if before[i].points != before[j].points {
			return before[i].points > before[j].points
		}
		return before[i].idx < before[j].idx
	})

	// Compute rank before: rank = position (1-indexed) when sorted descending
  // Alokasi slice
	rankBefore := make([]int, n)
	for pos, p := range before {
		rankBefore[p.idx] = pos + 1
	}

	// Create sorted list for after
	after := make([]pair, n)
	for i := 0; i < n; i++ {
		after[i] = pair{pointsAfter[i], i}
	}
  // Custom sort
	sort.Slice(after, func(i, j int) bool {
		if after[i].points != after[j].points {
			return after[i].points > after[j].points
		}
		return after[i].idx < after[j].idx
	})

	// Compute rank after
  // Alokasi slice
	rankAfter := make([]int, n)
	for pos, p := range after {
		rankAfter[p.idx] = pos + 1
	}

	// Difference
  // Alokasi slice
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rankBefore[i] - rankAfter[i]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", globalRankings([]int{10, 20, 30}, []int{15, 25, 35}))
	// Expected: [0, 0, 0] (same relative order)

	// Test case 2
	fmt.Println("Test 2:", globalRankings([]int{50, 40, 30, 20}, []int{45, 45, 35, 25}))
	// Expected: varying based on rank changes
}
```
