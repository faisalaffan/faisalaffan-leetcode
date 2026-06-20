# 2877 — Create A Dataframe From List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CreateADataframeFromList(studentData [][]int) DataFrame`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2877: Create a DataFrame from List
// https://leetcode.com/problems/create-a-dataframe-from-list/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we implement equivalent logic using a struct slice.

import "fmt"

func main() {
	// LeetCode name: createDataFrameFromList
	fmt.Println(CreateADataframeFromList([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}}))
	fmt.Println(CreateADataframeFromList([][]int{{5, 25}, {6, 30}}))
}

// type DataFrame represents the solution output.
type DataFrame []struct {
	StudentID int
	Age       int
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: createDataFrameFromList
func CreateADataframeFromList(studentData [][]int) DataFrame {
	df := make(DataFrame, len(studentData))
	for i, row := range studentData {
		df[i] = struct {
			StudentID int
			Age       int
		}{StudentID: row[0], Age: row[1]}
	}
	return df
}
```
