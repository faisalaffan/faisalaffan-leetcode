# 0646 — Maximum Length Of Pair Chain

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findLongestChain(pairs [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n) for sorting  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #646: Maximum Length of Pair Chain
// https://leetcode.com/problems/maximum-length-of-pair-chain/
// Difficulty: Medium
// Time: O(n log n) for sorting
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}

func findLongestChain(pairs [][]int) int {
  // Custom sort
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 0
	curEnd := -1 << 31

	for _, pair := range pairs {
		if pair[0] > curEnd {
			curEnd = pair[1]
			count++
		}
	}

	return count
}
```
