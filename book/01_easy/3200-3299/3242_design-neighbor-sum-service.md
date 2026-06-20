# 3242 — Design Neighbor Sum Service

## Deskripsi

**Soal:** [3242. Design Neighbor Sum Service](https://leetcode.com/problems/design-neighbor-sum-service/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2). Space: O(n^2).  
**Kompleksitas Ruang:** O(n^2).

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int][2]int)
  // Loop standar: indeks 0 sampai n-1
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
