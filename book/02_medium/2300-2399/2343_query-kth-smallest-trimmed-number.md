# 2343 — Query Kth Smallest Trimmed Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func smallestTrimmedNumbers(nums []string, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(m * n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2343: Query Kth Smallest Trimmed Number
// https://leetcode.com/problems/query-kth-smallest-trimmed-number/
// Difficulty: Medium
// Time: O(m * n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
  // Alokasi slice
	result := make([]int, len(queries))
	n := len(nums)

	type pair struct {
		val  string
		idx  int
	}

	for qi, q := range queries {
		k, trim := q[0], q[1]
		pairs := make([]pair, n)
		for i, s := range nums {
			pairs[i] = pair{s[len(s)-trim:], i}
		}
		sort.SliceStable(pairs, func(i, j int) bool {
			if pairs[i].val != pairs[j].val {
				return pairs[i].val < pairs[j].val
			}
			return pairs[i].idx < pairs[j].idx
		})
		result[qi] = pairs[k-1].idx
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}))
	// Expected: [2, 2, 1, 0]

	// Test case 2
	fmt.Println(smallestTrimmedNumbers([]string{"24", "37", "96", "04"}, [][]int{{2, 1}, {2, 2}}))
	// Expected: [3, 0]
}
```
