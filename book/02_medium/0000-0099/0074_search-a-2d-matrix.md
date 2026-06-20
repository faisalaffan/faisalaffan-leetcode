# 0074 — Search A 2D Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func searchMatrix(matrix [][]int, target int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log(m*n))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #74: Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/
// Difficulty: Medium

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

  // Binary search loop
	for left <= right {
		mid := left + (right-left)/2
		midVal := matrix[mid/n][mid%n]
		if midVal == target {
			return true
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	// Test case 1
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)) // true

	// Test case 2
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)) // false

	// Test case 3
	fmt.Println(searchMatrix([][]int{{1}}, 0)) // false
}

// Time: O(log(m*n)) | Space: O(1)
```
