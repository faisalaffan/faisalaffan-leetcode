# 1074 — Number Of Submatrices That Sum To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numSubmatrixSumTarget(matrix [][]int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1074: Number of Submatrices That Sum to Target
// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
// Difficulty: Hard
//
// Fix top row, expand bottom row, accumulate column-wise sums, then for each
// row-pair use prefix sum map (same as subarray sum equals target) across
// columns to count submatrices.

import "fmt"

func main() {
	matrix := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	fmt.Println(numSubmatrixSumTarget(matrix, 0))
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	ans := 0

	for top := 0; top < m; top++ {
  // Alokasi slice
		colSum := make([]int, n)
		for bottom := top; bottom < m; bottom++ {
			for c := 0; c < n; c++ {
				colSum[c] += matrix[bottom][c]
			}
			// Count subarrays in colSum that sum to target
			countMap := map[int]int{0: 1}
			prefix := 0
			for c := 0; c < n; c++ {
				prefix += colSum[c]
				ans += countMap[prefix-target]
				countMap[prefix]++
			}
		}
	}

	return ans
}
```
