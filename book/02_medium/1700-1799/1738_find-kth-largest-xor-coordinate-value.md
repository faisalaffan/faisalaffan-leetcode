# 1738 — Find Kth Largest Xor Coordinate Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func kthLargestValue(matrix [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** O(m * n * log(m*n)), Space: O(m * n)  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1738: Find Kth Largest XOR Coordinate Value
// https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/
// Difficulty: Medium
// Time: O(m * n * log(m*n)), Space: O(m * n)

import (
	"fmt"
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
  // Matriks 2D
	prefix := make([][]int, m)
	for i := 0; i < m; i++ {
		prefix[i] = make([]int, n)
	}

  // Alokasi slice
	values := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			if i > 0 {
				val ^= prefix[i-1][j]
			}
			if j > 0 {
				val ^= prefix[i][j-1]
			}
			if i > 0 && j > 0 {
				val ^= prefix[i-1][j-1]
			}
			prefix[i][j] = val
			values = append(values, val)
		}
	}

  // Custom sort
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values[k-1]
}

func main() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1)) // Expected: 7
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2)) // Expected: 5
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3)) // Expected: 4
}
```
