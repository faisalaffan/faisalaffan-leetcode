# 0778 — Swim In Rising Water

## Deskripsi

**Soal:** [0778. Swim In Rising Water](https://leetcode.com/problems/swim-in-rising-water/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Heap (priority queue), Dijkstra (lintasan terpendek)

> **Ide Kunci:** Dijkstra (min-heap)

## Solusi Go

```go
package main

// LeetCode #778: Swim in Rising Water
// https://leetcode.com/problems/swim-in-rising-water/
// Difficulty: Hard
//
// N x N grid where grid[i][j] = elevation. Water rises at time = t,
// you can swim only if elevation <= t. Find minimum time to go from
// (0,0) to (N-1,N-1).
//
// Approach: Dijkstra (min-heap)
// Treat each cell as a node with weight = elevation. The path cost is
// the maximum elevation on the path. Use a min-heap to always explore
// the cell with the smallest max-elevation-so-far.

import (
	"container/heap"
	"fmt"
)

type Item struct {
	r, c int
	cost int // max elevation encountered so far on path to this cell
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func main() {
	fmt.Println(swimInWater([][]int{{0, 2}, {1, 3}}))       // 3
	fmt.Println(swimInWater([][]int{{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6}})) // 16
	fmt.Println(swimInWater([][]int{{0, 3}, {2, 1}}))       // 2
}

func swimInWater(grid [][]int) int {
	n := len(grid)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, n)
  // Iterasi seluruh elemen
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	pq := make(PriorityQueue, 0, n*n)
	heap.Init(&pq)
	heap.Push(&pq, &Item{r: 0, c: 0, cost: grid[0][0]})
	visited[0][0] = true

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		if cur.r == n-1 && cur.c == n-1 {
			return cur.cost
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true
				newCost := cur.cost
				if grid[nr][nc] > newCost {
					newCost = grid[nr][nc]
				}
				heap.Push(&pq, &Item{r: nr, c: nc, cost: newCost})
			}
		}
	}

	return 0
}
```
