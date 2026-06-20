# 2577 — Minimum Time To Visit A Cell In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTime(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2577: Minimum Time to Visit a Cell in a Grid
// https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

// Item for min-heap priority queue
type Item struct {
	row, col, time int
}

type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// minimumTime finds minimum time to reach bottom-right using modified Dijkstra.
// We can "wait" by oscillating between two adjacent cells (2 time units per cycle).
//
// If both start neighbors require >1, impossible -> -1.
//
// Complexity: O(m*n*log(m*n)) time, O(m*n) space
func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Handle trivial 1x1 case
	if m == 1 && n == 1 {
		return 0
	}

	// If both neighbors of start are inaccessible at time 1
	hasRight := n > 1
	hasDown := m > 1
	if (!hasRight || grid[0][1] > 1) && (!hasDown || grid[1][0] > 1) {
		return -1
	}

  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
	}
	dist[0][0] = 0

	pq := &MinHeap{}
  // Masukkan elemen ke priority queue
	heap.Push(pq, Item{0, 0, 0})

	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(pq).(Item)
		if cur.time > dist[cur.row][cur.col] {
			continue
		}
		if cur.row == m-1 && cur.col == n-1 {
			return cur.time
		}

		for _, d := range dirs {
			nr, nc := cur.row+d[0], cur.col+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			nt := cur.time + 1
			if nt < grid[nr][nc] {
				diff := grid[nr][nc] - nt
				if diff%2 == 0 {
					nt = grid[nr][nc]
				} else {
					nt = grid[nr][nc] + 1
				}
			}

			if nt < dist[nr][nc] {
				dist[nr][nc] = nt
  // Masukkan elemen ke priority queue
				heap.Push(pq, Item{nr, nc, nt})
			}
		}
	}

	return -1
}

func main() {
	// Example from LeetCode
	grid1 := [][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}
	fmt.Println("Test 1: ->", minimumTime(grid1)) // 7

	// Additional test cases
	grid2 := [][]int{{0, 2, 4}, {3, 2, 1}, {1, 0, 4}}
	fmt.Println("Test 2: ->", minimumTime(grid2))

	grid3 := [][]int{{0, 2}, {3, 4}}
	fmt.Println("Test 3: blocked start ->", minimumTime(grid3))

	fmt.Println("Test 4: 1x1 ->", minimumTime([][]int{{0}})) // 0

	grid5 := [][]int{{0, 1}, {1, 0}}
	fmt.Println("Test 5: simple 2x2 ->", minimumTime(grid5))
}
```
