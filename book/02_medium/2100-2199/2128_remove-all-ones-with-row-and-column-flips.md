# 2128 — Remove All Ones With Row And Column Flips

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func removeOnes(grid [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2128: Remove All Ones With Row and Column Flips
// https://leetcode.com/problems/remove-all-ones-with-row-and-column-flips/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(1)

import "fmt"

func removeOnes(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		// Compare each row with first row: either equal or complementary
		same := true
		comp := true
		for j := 0; j < n; j++ {
			if grid[i][j] != grid[0][j] {
				same = false
			}
			if grid[i][j] != 1-grid[0][j] {
				comp = false
			}
		}
		if !same && !comp {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeOnes([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", removeOnes([][]int{{1, 1, 0}, {0, 0, 0}, {0, 0, 0}}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", removeOnes([][]int{{0, 0, 0}, {0, 0, 0}}))
	// Expected: true
}
```
