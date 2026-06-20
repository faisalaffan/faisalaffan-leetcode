# 2679 — Sum In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func matrixSum(nums [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(m * n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2679: Sum in a Matrix
// https://leetcode.com/problems/sum-in-a-matrix/
// Difficulty: Medium
// Time: O(m * n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	m := len(nums)
	if m == 0 {
		return 0
	}
	n := len(nums[0])

	// Sort each row
	for i := 0; i < m; i++ {
  // Sort O(n log n)
		sort.Ints(nums[i])
	}

	ans := 0
	for col := 0; col < n; col++ {
		maxVal := 0
		for row := 0; row < m; row++ {
			if nums[row][col] > maxVal {
				maxVal = nums[row][col]
			}
		}
		ans += maxVal
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", matrixSum([][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}))
	// Expected: 15

	// Test case 2
	fmt.Println("Test 2:", matrixSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// Expected: 18

	// Test case 3
	fmt.Println("Test 3:", matrixSum([][]int{{1}}))
	// Expected: 1
}
```
