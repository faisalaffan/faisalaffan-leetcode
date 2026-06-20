# 2880 — Select Data

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SelectData(df [][]int, studentID int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2880: Select Data
// https://leetcode.com/problems/select-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we select rows by student_id and return the age values.

import "fmt"

func main() {
	// LeetCode name: selectData
	fmt.Println(SelectData([][]int{{101, 15}, {102, 11}, {103, 11}}, 101)) // [15]
	fmt.Println(SelectData([][]int{{101, 20}, {102, 22}, {101, 21}}, 101)) // [20, 21]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: selectData
func SelectData(df [][]int, studentID int) []int {
	ages := []int{}
	for _, row := range df {
		if row[0] == studentID {
			ages = append(ages, row[1])
		}
	}
	return ages
}
```
