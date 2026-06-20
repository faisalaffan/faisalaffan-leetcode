# 1329 — Sort The Matrix Diagonally

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func diagonalSort(mat [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(m*n*log(min(m,n))) - sorting each diagonal  
**Kompleksitas Ruang:** O(m*n) for storing diagonal elements

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1329: Sort the Matrix Diagonally
// https://leetcode.com/problems/sort-the-matrix-diagonally/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(diagonalSort([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}))
	// [[1,1,1,1],[1,2,2,2],[1,2,3,3]]

	// Test case 2
	fmt.Println(diagonalSort([][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}}))

	// Test case 3
	fmt.Println(diagonalSort([][]int{{1}}))
	// [[1]]
}

// Time: O(m*n*log(min(m,n))) - sorting each diagonal
// Space: O(m*n) for storing diagonal elements
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	// Each diagonal starting from (i, 0) and (0, j)
	// Key insight: elements on same diagonal have same (i-j)

	// Group diagonals by (row - col) offset
  // Membuat map (HashMap) — pencarian O(1)
	diagonals := make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonals[i-j] = append(diagonals[i-j], mat[i][j])
		}
	}

	// Sort each diagonal
	for _, d := range diagonals {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(d)
	}

	// Place sorted values back
	// We need to track where we are in each diagonal
  // Membuat map (HashMap) — pencarian O(1)
	counters := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			offset := i - j
			mat[i][j] = diagonals[offset][counters[offset]]
			counters[offset]++
		}
	}

	return mat
}
```
