# 3242 — Design Neighbor Sum Service

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewNeighborSum(grid [][]int) NeighborSum
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2). Space: O(n^2).  
**Kompleksitas Ruang:** O(n^2).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3242: Design Neighbor Sum Service
// https://leetcode.com/problems/design-neighbor-sum-service/
// Difficulty: Easy

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	ns := NewNeighborSum(grid)
	fmt.Println(ns.AdjacentSum(1))
	fmt.Println(ns.AdjacentSum(4))
	fmt.Println(ns.DiagonalSum(4))
	fmt.Println(ns.DiagonalSum(8))
}

// NeighborSum provides sums of adjacent/diagonal elements for a given value in a grid.
type NeighborSum struct {
	grid [][]int
	pos  map[int][2]int // value -> [row, col]
}

// NewNeighborSum initializes a NeighborSum with the given grid.
// Time: O(n^2). Space: O(n^2).
func NewNeighborSum(grid [][]int) NeighborSum {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int][2]int)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			pos[grid[i][j]] = [2]int{i, j}
		}
	}
	return NeighborSum{grid: grid, pos: pos}
}

// AdjacentSum returns the sum of orthogonal neighbors of the cell containing value.
// Time: O(1). Space: O(1).
func (ns *NeighborSum) AdjacentSum(value int) int {
	p, ok := ns.pos[value]
	if !ok {
		return 0
	}
	r, c := p[0], p[1]
	sum := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(ns.grid) && nc >= 0 && nc < len(ns.grid[0]) {
			sum += ns.grid[nr][nc]
		}
	}
	return sum
}

// DiagonalSum returns the sum of diagonal neighbors of the cell containing value.
// Time: O(1). Space: O(1).
func (ns *NeighborSum) DiagonalSum(value int) int {
	p, ok := ns.pos[value]
	if !ok {
		return 0
	}
	r, c := p[0], p[1]
	sum := 0
	dirs := [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(ns.grid) && nc >= 0 && nc < len(ns.grid[0]) {
			sum += ns.grid[nr][nc]
		}
	}
	return sum
}
```
