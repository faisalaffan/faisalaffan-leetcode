# 2882 — Drop Duplicate Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func DropDuplicateRows(df [][]string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2882: Drop Duplicate Rows
// https://leetcode.com/problems/drop-duplicate-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows with duplicate emails.

import "fmt"

func main() {
	// LeetCode name: dropDuplicateEmails
	fmt.Println(DropDuplicateRows([][]string{{"1", "a@b.com"}, {"2", "c@d.com"}, {"3", "a@b.com"}}))
	// [[1 a@b.com] [2 c@d.com]]

	fmt.Println(DropDuplicateRows([][]string{{"1", "x@y.com"}, {"2", "x@y.com"}, {"3", "x@y.com"}}))
	// [[1 x@y.com]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropDuplicateEmails
func DropDuplicateRows(df [][]string) [][]string {
  // HashMap: O(1) lookup
	seen := make(map[string]bool)
	result := [][]string{}
	for _, row := range df {
		email := row[1]
		if !seen[email] {
			seen[email] = true
			result = append(result, row)
		}
	}
	return result
}
```
