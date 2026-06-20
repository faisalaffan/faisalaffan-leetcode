# 2884 — Modify Columns

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ModifyColumns(df [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2884: Modify Columns
// https://leetcode.com/problems/modify-columns/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we modify a column by multiplying salary by 2.

import "fmt"

func main() {
	// LeetCode name: modifySalaryColumn
	fmt.Println(ModifyColumns([][]int{{1, 100}, {2, 200}, {3, 300}}))
	// [[1 200] [2 400] [3 600]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: modifySalaryColumn
func ModifyColumns(df [][]int) [][]int {
  // Matriks 2D
	result := make([][]int, len(df))
	for i, row := range df {
		result[i] = []int{row[0], row[1] * 2}
	}
	return result
}
```
