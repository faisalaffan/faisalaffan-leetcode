# 0308 — Range Sum Query 2D Mutable

## Deskripsi

**Soal:** [0308. Range Sum Query 2D Mutable](https://leetcode.com/problems/range-sum-query-2d-mutable/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log m * log n) for update/sum, Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(matrix [][]int) NumMatrix`

## Solusi Go

```go
package main

// LeetCode #308: Range Sum Query 2D - Mutable
// https://leetcode.com/problems/range-sum-query-2d-mutable/
// Difficulty: Medium [Paid]
// Time: O(log m * log n) for update/sum, Space: O(m*n)

import "fmt"

type NumMatrix struct {
	matrix [][]int
	bit    [][]int
	rows   int
	cols   int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	rows, cols := len(matrix), len(matrix[0])
  // Membuat slice 2D untuk DP/tabel
	bit := make([][]int, rows+1)
  // Iterasi seluruh elemen
	for i := range bit {
		bit[i] = make([]int, cols+1)
	}

	nm := NumMatrix{matrix, bit, rows, cols}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			nm.add(r+1, c+1, matrix[r][c])
		}
	}
	return nm
}

func (this *NumMatrix) add(row, col, val int) {
	for r := row; r <= this.rows; r += r & -r {
		for c := col; c <= this.cols; c += c & -c {
			this.bit[r][c] += val
		}
	}
}

func (this *NumMatrix) sum(row, col int) int {
	res := 0
	for r := row; r > 0; r -= r & -r {
		for c := col; c > 0; c -= c & -c {
			res += this.bit[r][c]
		}
	}
	return res
}

func (this *NumMatrix) Update(row int, col int, val int) {
	diff := val - this.matrix[row][col]
	this.matrix[row][col] = val
	this.add(row+1, col+1, diff)
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return this.sum(row2+1, col2+1) - this.sum(row1, col2+1) - this.sum(row2+1, col1) + this.sum(row1, col1)
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
	nm.Update(3, 2, 2)
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
}
```
