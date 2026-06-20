# 3112 — Minimum Time To Visit Disappearing Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func minimumTime(n int, edges [][]int, disappear []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O((n + m) log n)  |  **Ruang:** O(n + m)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3112: Minimum Time to Visit Disappearing Nodes
// https://leetcode.com/problems/minimum-time-to-visit-disappearing-nodes/
// Difficulty: Medium
// Time: O((n + m) log n) | Space: O(n + m)

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, w int
}

type Item struct {
	node, dist int
	idx        int
}

type PQ []*Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}
func (pq *PQ) Push(x any) { *pq = append(*pq, x.(*Item)) }
func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func minimumTime(n int, edges [][]int, disappear []int) []int {
  // Matriks 2D
	graph := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		graph[v] = append(graph[v], Edge{u, w})
	}

  // Alokasi slice
	dist := make([]int, n)
  // Range loop
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0

	pq := &PQ{}
	heap.Init(pq)
  // Push ke priority queue
	heap.Push(pq, &Item{node: 0, dist: 0})

	for pq.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(pq).(*Item)
		if cur.dist > dist[cur.node] {
			continue
		}

		for _, e := range graph[cur.node] {
			nd := cur.dist + e.w
			if nd < disappear[e.to] && nd < dist[e.to] {
				dist[e.to] = nd
  // Push ke priority queue
				heap.Push(pq, &Item{node: e.to, dist: nd})
			}
		}
	}

  // Alokasi slice
	ans := make([]int, n)
  // Range loop
	for i := range ans {
		if dist[i] == math.MaxInt32 {
			ans[i] = -1
		} else {
			ans[i] = dist[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumTime(3, [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}, []int{1, 1, 5}))
	fmt.Println(minimumTime(3, [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}, []int{1, 3, 5}))
	fmt.Println(minimumTime(2, [][]int{{0, 1, 1}}, []int{1, 1}))
}
```
