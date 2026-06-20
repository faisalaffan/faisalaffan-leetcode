# 2617 — Minimum Number Of Visited Cells In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func newDSU(n int) *dsu
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2617: Minimum Number of Visited Cells in a Grid
// https://leetcode.com/problems/minimum-number-of-visited-cells-in-a-grid/
// Difficulty: Hard
//
// From (i,j) you can move to (i, j+k) right or (i+k, j) down, 1 <= k <= grid[i][j].
// Find minimum cells visited from (0,0) to (m-1,n-1). Return -1 if impossible.
// BFS with DSU skip-list to avoid O(n^2) per cell.

import "fmt"

func main() {
	// Example 1: grid = [[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]] -> 4
	fmt.Println(minimumVisitedCells([][]int{{3, 4, 2, 1}, {4, 2, 3, 1}, {2, 1, 0, 0}, {2, 4, 0, 0}}))
	// Example 2: single cell
	fmt.Println(minimumVisitedCells([][]int{{0}}))
	// Example 3: no path
	fmt.Println(minimumVisitedCells([][]int{{1, 0}, {0, 1}}))
}

// DSU with path compression, tracks next unvisited index
type dsu struct {
	p []int
}

func newDSU(n int) *dsu {
  // Alokasi slice integer
	p := make([]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range p {
		p[i] = i
	}
	return &dsu{p}
}

func (d *dsu) find(x int) int {
	if d.p[x] != x {
		d.p[x] = d.find(d.p[x])
	}
	return d.p[x]
}

// mark x as visited by unioning with x+1
func (d *dsu) mark(x int) {
	d.p[x] = d.find(x + 1)
}

func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dist stores steps from (0,0), -1 = unvisited
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// DSU per row for skipping columns, per column for skipping rows
	rowDSU := make([]*dsu, m)
	colDSU := make([]*dsu, n)
	for i := 0; i < m; i++ {
		rowDSU[i] = newDSU(n)
	}
	for j := 0; j < n; j++ {
		colDSU[j] = newDSU(m)
	}

	dist[0][0] = 1
	rowDSU[0].mark(0)
	colDSU[0].mark(0)

	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		i, j := cur[0], cur[1]
		d := dist[i][j]
		k := grid[i][j]

		if k <= 0 {
			continue
		}

		// Move right: next column in row i
		rightBound := j + k
		if rightBound > n-1 {
			rightBound = n - 1
		}
		for c := rowDSU[i].find(j + 1); c <= rightBound; c = rowDSU[i].find(c + 1) {
			if dist[i][c] == -1 {
				dist[i][c] = d + 1
				q = append(q, [2]int{i, c})
			}
			rowDSU[i].mark(c)
			colDSU[c].mark(i)
		}

		// Move down: next row in column j
		downBound := i + k
		if downBound > m-1 {
			downBound = m - 1
		}
		for r := colDSU[j].find(i + 1); r <= downBound; r = colDSU[j].find(r + 1) {
			if dist[r][j] == -1 {
				dist[r][j] = d + 1
				q = append(q, [2]int{r, j})
			}
			colDSU[j].mark(r)
			rowDSU[r].mark(j)
		}
	}

	return dist[m-1][n-1]
}
```
