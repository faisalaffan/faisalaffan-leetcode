# 2737 — Find The Closest Marked Node

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func FindTheClosestMarkedNode(n int, edges [][]int, marked []int, start int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS, Heap

**Waktu:** O((V+E) log V)  |  **Ruang:** O(V+E)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2737: Find the Closest Marked Node
// https://leetcode.com/problems/find-the-closest-marked-node/
// Difficulty: Medium [Paid]
// Time: O((V+E) log V) | Space: O(V+E)

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
	index      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) { n := len(*pq); item := x.(*Item); item.index = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{}   { old := *pq; n := len(old); item := old[n-1]; old[n-1] = nil; item.index = -1; *pq = old[:n-1]; return item }

func FindTheClosestMarkedNode(n int, edges [][]int, marked []int, start int) int {
  // Matriks 2D
	graph := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
	}

  // Alokasi slice
	dist := make([]int, n)
  // Range loop
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
  // Push ke priority queue
	heap.Push(pq, &Item{node: start, dist: 0})

	for pq.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(pq).(*Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.dist + e.weight; nd < dist[e.to] {
				dist[e.to] = nd
  // Push ke priority queue
				heap.Push(pq, &Item{node: e.to, dist: nd})
			}
		}
	}

	best := math.MaxInt32
	for _, m := range marked {
		if dist[m] < best {
			best = dist[m]
		}
	}
	if best == math.MaxInt32 {
		return -1
	}
	return best
}

func main() {
	fmt.Println(FindTheClosestMarkedNode(4, [][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}}, []int{2, 3}, 0))
	fmt.Println(FindTheClosestMarkedNode(3, [][]int{{0, 1, 5}}, []int{2}, 0))
}
```
