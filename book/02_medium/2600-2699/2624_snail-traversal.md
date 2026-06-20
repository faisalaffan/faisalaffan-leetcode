# 2624 — Snail Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func snailTraversal(arr []int, rowsCount int, colsCount int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n*m)  |  **Ruang:** O(n*m)


## 💻 Solusi Go

```go
package main

// LeetCode #2624: Snail Traversal
// https://leetcode.com/problems/snail-traversal/
// Difficulty: Medium
// Time: O(n*m) | Space: O(n*m)

import "fmt"

func snailTraversal(arr []int, rowsCount int, colsCount int) [][]int {
	n := len(arr)
	if rowsCount*colsCount != n {
		return [][]int{}
	}

  // Matriks 2D
	result := make([][]int, rowsCount)
  // Range loop
	for i := range result {
		result[i] = make([]int, colsCount)
	}

	idx := 0
	for col := 0; col < colsCount; col++ {
		if col%2 == 0 {
			// Top to bottom
			for row := 0; row < rowsCount; row++ {
				result[row][col] = arr[idx]
				idx++
			}
		} else {
			// Bottom to top
			for row := rowsCount - 1; row >= 0; row-- {
				result[row][col] = arr[idx]
				idx++
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", snailTraversal([]int{1, 2, 3, 4}, 2, 2))
	// Expected: [[1,4],[2,3]] (snail fill)

	// Test case 2
	fmt.Println("Test 2:", snailTraversal([]int{1, 2, 3, 4, 5, 6}, 2, 3))
	// Expected: [[1,4,5],[2,3,6]]

	// Test case 3: invalid dimensions
	fmt.Println("Test 3:", snailTraversal([]int{1, 2, 3}, 2, 2))
	// Expected: []
}
```
