# 2883 — Drop Missing Data

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func DropMissingData(df [][]string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2883: Drop Missing Data
// https://leetcode.com/problems/drop-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows where the name column is empty.

import "fmt"

func main() {
	// LeetCode name: dropMissingData
	fmt.Println(DropMissingData([][]string{{"1", "Alice", "15"}, {"2", "", "11"}, {"3", "Bob", "12"}}))
	// [[1 Alice 15] [3 Bob 12]]

	fmt.Println(DropMissingData([][]string{{"1", "", "10"}}))
	// []
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropMissingData
func DropMissingData(df [][]string) [][]string {
	result := [][]string{}
	for _, row := range df {
		if row[1] != "" {
			result = append(result, row)
		}
	}
	return result
}
```
