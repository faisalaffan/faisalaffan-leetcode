# 0251 — Flatten 2D Vector

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func Constructor(vec [][]int) Vector2D`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1) amortized per next/hasNext, Space: O(1) excluding input  |  **Ruang:** O(1) excluding input


## 💻 Solusi Go

```go
package main

// LeetCode #251: Flatten 2D Vector
// https://leetcode.com/problems/flatten-2d-vector/
// Difficulty: Medium [Paid]
// Time: O(1) amortized per next/hasNext, Space: O(1) excluding input

import "fmt"

type Vector2D struct {
	vec    [][]int
	row    int
	col    int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{vec, 0, 0}
}

func (this *Vector2D) advance() {
	for this.row < len(this.vec) && this.col >= len(this.vec[this.row]) {
		this.row++
		this.col = 0
	}
}

func (this *Vector2D) Next() int {
	this.advance()
	val := this.vec[this.row][this.col]
	this.col++
	return val
}

func (this *Vector2D) HasNext() bool {
	this.advance()
	return this.row < len(this.vec)
}

func main() {
	iter := Constructor([][]int{{1, 2}, {3}, {4, 5, 6}})
	for iter.HasNext() {
		fmt.Print(iter.Next(), " ")
	}
	fmt.Println()

	iter2 := Constructor([][]int{{}, {1}, {}})
	for iter2.HasNext() {
		fmt.Print(iter2.Next(), " ")
	}
	fmt.Println()

	iter3 := Constructor([][]int{{}})
	fmt.Println(iter3.HasNext())
}
```
