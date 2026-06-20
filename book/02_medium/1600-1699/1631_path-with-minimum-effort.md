# 1631 — Path With Minimum Effort

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumEffortPath(heights [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** O(R*C log(R*C)), Space: O(R*C)  
**Kompleksitas Ruang:** O(R*C)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1631: Path With Minimum Effort
// https://leetcode.com/problems/path-with-minimum-effort/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	fmt.Println(MinimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
	fmt.Println(MinimumEffortPath([][]int{{1, 10, 6, 7, 9, 10, 4, 9}}))
}

type Point struct {
	x, y, effort int
	index        int
}

type EffortPQ []*Point

func (pq EffortPQ) Len() int           { return len(pq) }
func (pq EffortPQ) Less(i, j int) bool { return pq[i].effort < pq[j].effort }
func (pq EffortPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *EffortPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Point)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *EffortPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func MinimumEffortPath(heights [][]int) int {
	// Time: O(R*C log(R*C)), Space: O(R*C)
	// Dijkstra-like: track minimum effort to reach each cell
	rows, cols := len(heights), len(heights[0])
	if rows == 0 || cols == 0 {
		return 0
	}

  // Membuat matriks/slice 2D untuk DP
	effort := make([][]int, rows)
	for i := 0; i < rows; i++ {
		effort[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			effort[i][j] = math.MaxInt32
		}
	}
	effort[0][0] = 0

	pq := &EffortPQ{}
  // Masukkan elemen ke priority queue
	heap.Push(pq, &Point{x: 0, y: 0, effort: 0})

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		p := heap.Pop(pq).(*Point)
		if p.effort > effort[p.x][p.y] {
			continue
		}
		if p.x == rows-1 && p.y == cols-1 {
			return p.effort
		}

		for _, d := range dirs {
			nx, ny := p.x+d[0], p.y+d[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
				continue
			}
			diff := heights[p.x][p.y] - heights[nx][ny]
			if diff < 0 {
				diff = -diff
			}
			newEffort := p.effort
			if diff > newEffort {
				newEffort = diff
			}
			if newEffort < effort[nx][ny] {
				effort[nx][ny] = newEffort
  // Masukkan elemen ke priority queue
				heap.Push(pq, &Point{x: nx, y: ny, effort: newEffort})
			}
		}
	}

	return 0
}
```
