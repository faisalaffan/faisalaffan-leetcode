# 3538 — Merge Operations For Minimum Travel Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func minimumTravelTime(n int, edges [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS, DP, Heap, Dijkstra

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3538: Merge Operations for Minimum Travel Time
// https://leetcode.com/problems/merge-operations-for-minimum-travel-time/
// Difficulty: Hard
//
// Given a graph with travel times, merge nodes to minimize travel time
// between start and end. Each merge combines two adjacent nodes.
//
// Approach: DP on intervals or use Dijkstra with state compression.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minimumTravelTime(4, [][]int{{0, 1, 2}, {0, 2, 5}, {2, 3, 1}, {1, 3, 3}}))
	// Example 2
	fmt.Println(minimumTravelTime(3, [][]int{{0, 1, 1}, {1, 2, 2}}))
	// Edge: single edge
	fmt.Println(minimumTravelTime(2, [][]int{{0, 1, 5}}))
}

type Item struct {
	node int
	dist int64
	idx  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].idx = i; pq[j].idx = j }
func (pq *PriorityQueue) Push(x interface{}) { n := len(*pq); item := x.(*Item); item.idx = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{} { old := *pq; n := len(old); item := old[n-1]; item.idx = -1; *pq = old[:n-1]; return item }

func minimumTravelTime(n int, edges [][]int) int64 {
  // Matriks 2D
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// Dijkstra from 0 to n-1
  // Alokasi slice
	dist := make([]int64, n)
  // Range loop
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
  // Push ke priority queue
	heap.Push(pq, &Item{node: 0, dist: 0})

	for pq.Len() > 0 {
  // Pop dari priority queue
		item := heap.Pop(pq).(*Item)
		u := item.node
		if item.dist > dist[u] {
			continue
		}
		if u == n-1 {
			return item.dist
		}
		for _, edge := range adj[u] {
			v, w := edge[0], int64(edge[1])
			if nd := item.dist + w; nd < dist[v] {
				dist[v] = nd
  // Push ke priority queue
				heap.Push(pq, &Item{node: v, dist: nd})
			}
		}
	}

	return -1
}
```
