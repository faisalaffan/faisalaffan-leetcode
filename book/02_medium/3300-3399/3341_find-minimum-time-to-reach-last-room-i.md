# 3341 — Find Minimum Time To Reach Last Room I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minTimeToReach(moveTime [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O(m * n * log(m * n)) Space: O(m * n)  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3341: Find Minimum Time to Reach Last Room I
// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/
// Difficulty: Medium
// Time: O(m * n * log(m * n)) Space: O(m * n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}})) // 6
	fmt.Println(minTimeToReach([][]int{{0, 0, 0}, {0, 0, 0}})) // 2
}

type State struct {
	time int
	r    int
	c    int
}

type MinHeap []State

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
  // Matriks 2D
	dist := make([][]int, m)
  // Range loop
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
	}
	dist[0][0] = 0

	h := &MinHeap{}
	heap.Init(h)
  // Push ke priority queue
	heap.Push(h, State{0, 0, 0})
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(h).(State)
		if cur.time > dist[cur.r][cur.c] {
			continue
		}
		if cur.r == m-1 && cur.c == n-1 {
			return cur.time
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				wait := 0
				if cur.time < moveTime[nr][nc] {
					wait = moveTime[nr][nc] - cur.time
				}
				nt := cur.time + 1 + wait
				if nt < dist[nr][nc] {
					dist[nr][nc] = nt
  // Push ke priority queue
					heap.Push(h, State{nt, nr, nc})
				}
			}
		}
	}
	return -1
}
```
