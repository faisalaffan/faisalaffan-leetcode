# 1975 — Maximum Matrix Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaxMatrixSum(matrix [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1975: Maximum Matrix Sum
// https://leetcode.com/problems/maximum-matrix-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(MaxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
}

// Time: O(m*n), Space: O(1)
func MaxMatrixSum(matrix [][]int) int64 {
	total := int64(0)
	negCount := 0
	minAbs := int64(1 << 31)

	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				negCount++
			}
			abs := int64(val)
			if abs < 0 {
				abs = -abs
			}
			total += abs
			if abs < minAbs {
				minAbs = abs
			}
		}
	}

	if negCount%2 == 1 {
		total -= 2 * minAbs
	}
	return total
}
```
