# 2888 — Reshape Data Concatenate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ReshapeDataConcatenate(df1, df2 [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n+m)  |  **Ruang:** O(n+m)


## 💻 Solusi Go

```go
package main

// LeetCode #2888: Reshape Data: Concatenate
// https://leetcode.com/problems/reshape-data-concatenate/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we concatenate two dataframes vertically.

import "fmt"

func main() {
	// LeetCode name: concatenateDataFrames
	fmt.Println(ReshapeDataConcatenate([][]int{{1, 15}, {2, 11}}, [][]int{{3, 12}, {4, 14}}))
	// [[1 15] [2 11] [3 12] [4 14]]
}

// Time: O(n+m) | Space: O(n+m)
// LeetCode submission name: concatenateDataFrames
func ReshapeDataConcatenate(df1, df2 [][]int) [][]int {
  // Matriks 2D
	result := make([][]int, 0, len(df1)+len(df2))
	result = append(result, df1...)
	result = append(result, df2...)
	return result
}
```
