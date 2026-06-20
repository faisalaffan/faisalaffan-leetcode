# 3142 — Check If Grid Satisfies Conditions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CheckIfGridSatisfiesConditions(grid [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3142: Check if Grid Satisfies Conditions
// https://leetcode.com/problems/check-if-grid-satisfies-conditions/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: satisfiesConditions
	fmt.Println(CheckIfGridSatisfiesConditions([][]int{{1, 0, 2}, {1, 0, 2}})) // true
	fmt.Println(CheckIfGridSatisfiesConditions([][]int{{1, 1, 1}, {0, 0, 0}})) // false
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: satisfiesConditions
func CheckIfGridSatisfiesConditions(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Below must be equal
			if i+1 < m && grid[i][j] != grid[i+1][j] {
				return false
			}
			// Right must be different
			if j+1 < n && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}
```
