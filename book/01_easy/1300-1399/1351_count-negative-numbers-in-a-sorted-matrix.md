# 1351 — Count Negative Numbers In A Sorted Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func countNegatives(grid [][]int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m + n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1351: Count Negative Numbers in a Sorted Matrix
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
// Difficulty: Easy
//
// LeetCode submission: func countNegatives(grid [][]int) int

import "fmt"

func main() {
	grid1 := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(CountNegativeNumbersInASortedMatrix(grid1)) // 8

	grid2 := [][]int{
		{3, 2},
		{1, 0},
	}
	fmt.Println(CountNegativeNumbersInASortedMatrix(grid2)) // 0
}

// Time: O(m + n), Space: O(1)
func CountNegativeNumbersInASortedMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	row, col := 0, n-1
	for row < m && col >= 0 {
		if grid[row][col] < 0 {
			count += m - row
			col--
		} else {
			row++
		}
	}
	return count
}
```
