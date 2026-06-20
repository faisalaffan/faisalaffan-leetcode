# 0304 — Range Sum Query 2D Immutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(matrix [][]int) NumMatrix
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(m*n) for init, O(1) for sum, Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #304: Range Sum Query 2D - Immutable
// https://leetcode.com/problems/range-sum-query-2d-immutable/
// Difficulty: Medium
// Time: O(m*n) for init, O(1) for sum, Space: O(m*n)

import "fmt"

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	rows, cols := len(matrix), len(matrix[0])
  // Membuat matriks/slice 2D untuk DP
	prefix := make([][]int, rows+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range prefix {
		prefix[i] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			prefix[r+1][c+1] = matrix[r][c] + prefix[r][c+1] + prefix[r+1][c] - prefix[r][c]
		}
	}

	return NumMatrix{prefix}
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return this.prefix[row2+1][col2+1] - this.prefix[row1][col2+1] - this.prefix[row2+1][col1] + this.prefix[row1][col1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	nm := Constructor(matrix)
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
	fmt.Println(nm.SumRegion(1, 1, 2, 2))
	fmt.Println(nm.SumRegion(1, 2, 2, 4))
}
```
