# 2352 — Equal Row And Column Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func equalPairs(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2352: Equal Row and Column Pairs
// https://leetcode.com/problems/equal-row-and-column-pairs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import (
	"fmt"
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
  // HashMap: O(1) lookup
	rowMap := make(map[string]int)

	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := 0; j < n; j++ {
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(grid[i][j]))
		}
		rowMap[sb.String()]++
	}

	count := 0
	for j := 0; j < n; j++ {
		var sb strings.Builder
		for i := 0; i < n; i++ {
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(grid[i][j]))
		}
		count += rowMap[sb.String()]
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	// Expected: 1

	// Test case 2
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
	// Expected: 3

	// Test case 3
	fmt.Println(equalPairs([][]int{{11, 1}, {1, 11}}))
	// Expected: 2
}
```
