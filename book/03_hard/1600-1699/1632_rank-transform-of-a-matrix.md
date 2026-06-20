# 1632 — Rank Transform Of A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewUnionFind(n int) *UnionFind
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1632: Rank Transform of a Matrix
// https://leetcode.com/problems/rank-transform-of-a-matrix/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
  // Alokasi slice integer
	parent := make([]int, n)
  // Alokasi slice integer
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	xr, yr := uf.Find(x), uf.Find(y)
	if xr == yr {
		return
	}
	if uf.rank[xr] < uf.rank[yr] {
		uf.parent[xr] = yr
	} else if uf.rank[xr] > uf.rank[yr] {
		uf.parent[yr] = xr
	} else {
		uf.parent[yr] = xr
		uf.rank[xr]++
	}
}

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	// Group cells by their value
  // Membuat map (HashMap) — pencarian O(1)
	valToCells := make(map[int][][2]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			valToCells[val] = append(valToCells[val], [2]int{i, j})
		}
	}

	// Sort unique values
  // Alokasi slice integer
	values := make([]int, 0, len(valToCells))
	for v := range valToCells {
		values = append(values, v)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(values)

	// Track the current max rank for each row and column
  // Alokasi slice integer
	rowMax := make([]int, m)
  // Alokasi slice integer
	colMax := make([]int, n)

	for _, val := range values {
		cells := valToCells[val]

		// Union cells that are in the same row or column (same value)
		uf := NewUnionFind(len(cells))
  // Membuat map (HashMap) — pencarian O(1)
		rowMap := make(map[int]int) // row -> first cell index
  // Membuat map (HashMap) — pencarian O(1)
		colMap := make(map[int]int) // col -> first cell index

		for idx, cell := range cells {
			r, c := cell[0], cell[1]
			if first, ok := rowMap[r]; ok {
				uf.Union(idx, first)
			} else {
				rowMap[r] = idx
			}
			if first, ok := colMap[c]; ok {
				uf.Union(idx, first)
			} else {
				colMap[c] = idx
			}
		}

		// Group cells by their root (connected components)
  // Membuat map (HashMap) — pencarian O(1)
		groups := make(map[int][]int)
		for idx := range cells {
			root := uf.Find(idx)
			groups[root] = append(groups[root], idx)
		}

		// For each connected component, compute rank = max(rowMax, colMax) + 1
		for _, group := range groups {
			maxRank := 0
			for _, idx := range group {
				r, c := cells[idx][0], cells[idx][1]
				if rowMax[r] > maxRank {
					maxRank = rowMax[r]
				}
				if colMax[c] > maxRank {
					maxRank = colMax[c]
				}
			}
			rank := maxRank + 1
			for _, idx := range group {
				r, c := cells[idx][0], cells[idx][1]
				result[r][c] = rank
				rowMax[r] = rank
				colMax[c] = rank
			}
		}
	}

	return result
}

func main() {
	// Test case 1: [[1,2],[3,4]] -> [[1,2],[2,3]]
	matrix := [][]int{{1, 2}, {3, 4}}
	result := matrixRankTransform(matrix)
	fmt.Printf("matrix=%v -> %v (expected [[1,2],[2,3]])\n", matrix, result)

	// Test case 2: [[7,7],[7,7]] -> [[1,1],[1,1]]
	matrix2 := [][]int{{7, 7}, {7, 7}}
	result2 := matrixRankTransform(matrix2)
	fmt.Printf("matrix=%v -> %v (expected [[1,1],[1,1]])\n", matrix2, result2)

	// Test case 3: [[20,-21,14],[-19,4,19],[22,-47,24],[-19,4,19]]
	matrix3 := [][]int{{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19}}
	result3 := matrixRankTransform(matrix3)
	fmt.Printf("matrix=%v -> %v\n", matrix3, result3)
}
```
