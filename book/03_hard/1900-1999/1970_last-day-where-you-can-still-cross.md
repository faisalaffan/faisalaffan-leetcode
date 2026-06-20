# 1970 — Last Day Where You Can Still Cross

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewDSU(n int) *DSU
```

> **💡 Hint:** Binary search + Union-Find (DSU)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1970: Last Day Where You Can Still Cross
// https://leetcode.com/problems/last-day-where-you-can-still-cross/
// Difficulty: Hard
// Approach: Binary search + Union-Find (DSU)
// For each day mid, build DSU with cells up to day mid as water.
// Check if there's a path from top to bottom through land cells.
// Virtual nodes: row*col = top, row*col+1 = bottom.

import "fmt"

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
  // Alokasi slice integer
	p := make([]int, n)
  // Alokasi slice integer
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	xr, yr := d.Find(x), d.Find(y)
	if xr == yr {
		return
	}
	if d.rank[xr] < d.rank[yr] {
		xr, yr = yr, xr
	}
	d.parent[yr] = xr
	if d.rank[xr] == d.rank[yr] {
		d.rank[xr]++
	}
}

func latestDayToCross(row, col int, cells [][]int) int {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n := row * col
	top, bottom := n, n+1

	canCross := func(day int) bool {
		dsu := NewDSU(n + 2)
  // Membuat matriks/slice 2D untuk DP
		land := make([][]bool, row)
		for i := 0; i < row; i++ {
			land[i] = make([]bool, col)
			for j := 0; j < col; j++ {
				land[i][j] = true
			}
		}
		// Flood first 'day' cells
		for d := 0; d < day; d++ {
			r, c := cells[d][0]-1, cells[d][1]-1
			land[r][c] = false
		}

		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				if !land[r][c] {
					continue
				}
				idx := r*col + c
				if r == 0 {
					dsu.Union(idx, top)
				}
				if r == row-1 {
					dsu.Union(idx, bottom)
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < row && nc >= 0 && nc < col && land[nr][nc] {
						dsu.Union(idx, nr*col+nc)
					}
				}
			}
		}
		return dsu.Find(top) == dsu.Find(bottom)
	}

	lo, hi := 1, row*col
	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if canCross(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func main() {
	// Example: row=2, col=2, cells=[[1,1],[2,1],[1,2],[2,2]] -> 2
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}))

	// Additional tests
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(latestDayToCross(3, 3, [][]int{{1, 1}, {2, 1}, {3, 1}, {1, 2}, {2, 2}, {3, 2}, {1, 3}, {2, 3}, {3, 3}}))
}
```
