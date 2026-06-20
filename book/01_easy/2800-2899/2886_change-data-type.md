# 2886 — Change Data Type

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ChangeDataType(df [][]float64) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2886: Change Data Type
// https://leetcode.com/problems/change-data-type/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we convert the grade column from float64 to int (truncation).

import "fmt"

func main() {
	// LeetCode name: changeDataType
	// Input: [student_id, grade (float)]
	fmt.Println(ChangeDataType([][]float64{{1, 3.5}, {2, 4.2}, {3, 2.8}}))
	// [[1 3] [2 4] [3 2]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: changeDataType
func ChangeDataType(df [][]float64) [][]int {
  // Matriks 2D
	result := make([][]int, len(df))
	for i, row := range df {
		result[i] = []int{int(row[0]), int(row[1])}
	}
	return result
}
```
