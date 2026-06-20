# 0304 — Range Sum Query 2D Immutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func Constructor(matrix [][]int) NumMatrix`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(m*n) for init, O(1) for sum, Space: O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

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
  // Matriks 2D
	prefix := make([][]int, rows+1)
  // Range loop
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
