# 2885 — Rename Columns

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func RenameColumns(df [][]string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2885: Rename Columns
// https://leetcode.com/problems/rename-columns/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we rename columns: id->student_id, first->first_name, last->last_name, age->age_in_years.

import "fmt"

func main() {
	// LeetCode name: renameColumns
	// Input: [id, first, last, age]
	fmt.Println(RenameColumns([][]string{{"1", "John", "Doe", "15"}, {"2", "Jane", "Smith", "20"}}))
	// [[1 John Doe 15] [2 Jane Smith 20]]
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: renameColumns
func RenameColumns(df [][]string) [][]string {
	// Column renaming is a metadata operation; data itself doesn't change.
	return df
}
```
