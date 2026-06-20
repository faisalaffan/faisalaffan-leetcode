# 2881 — Create A New Column

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CreateANewColumn(df [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2881: Create a New Column
// https://leetcode.com/problems/create-a-new-column/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we add a "grade" column computed from existing data.

import "fmt"

func main() {
	// LeetCode name: createBonusColumn
	fmt.Println(CreateANewColumn([][]int{{101, 15}, {102, 11}, {103, 20}}))
	// [[101 15 30] [102 11 22] [103 20 40]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: createBonusColumn
func CreateANewColumn(df [][]int) [][]int {
  // Matriks 2D
	result := make([][]int, len(df))
	for i, row := range df {
		// bonus = salary * 2
		result[i] = []int{row[0], row[1], row[1] * 2}
	}
	return result
}
```
