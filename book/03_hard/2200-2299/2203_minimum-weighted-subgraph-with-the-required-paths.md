# 2203 — Minimum Weighted Subgraph With The Required Paths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func dijkstra(n int, graph [][]Edge, src int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2203: Minimum Weighted Subgraph with the Required Paths
// https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/
// Difficulty: Hard
//
// 3 Dijkstras: Compute shortest distances from src1, src2, and to dest (reverse graph).
// Answer = min over all nodes v of dist(src1=>v) + dist(src2=>v) + dist(v=>dest).

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	n := 6
	edges := [][]int{
		{0, 2, 2},
		{0, 5, 6},
		{1, 0, 3},
		{1, 4, 5},
		{2, 1, 1},
		{2, 3, 3},
		{2, 3, 4},
		{3, 4, 2},
		{4, 5, 1},
	}
	src1, src2, dest := 0, 1, 5
	// Expected: 9
	fmt.Println(minimumWeight(n, edges, src1, src2, dest))

	// Single node
	fmt.Println(minimumWeight(1, [][]int{}, 0, 0, 0))
}

type Edge struct {
	to, w int
}

type Item struct {
	node, dist int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func dijkstra(n int, graph [][]Edge, src int) []int {
  // Alokasi slice integer
	dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[src] = 0
	h := &MinHeap{{src, 0}}
	heap.Init(h)

	for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(h).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			nd := cur.dist + e.w
			if nd < dist[e.to] {
				dist[e.to] = nd
  // Masukkan elemen ke priority queue
				heap.Push(h, Item{e.to, nd})
			}
		}
	}
	return dist
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]Edge, n)
  // Membuat matriks/slice 2D untuk DP
	rgraph := make([][]Edge, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		rgraph[v] = append(rgraph[v], Edge{u, w})
	}

	d1 := dijkstra(n, graph, src1)
	d2 := dijkstra(n, graph, src2)
	dd := dijkstra(n, rgraph, dest)

	ans := math.MaxInt64
	for v := 0; v < n; v++ {
		if d1[v] == math.MaxInt64 || d2[v] == math.MaxInt64 || dd[v] == math.MaxInt64 {
			continue
		}
		total := d1[v] + d2[v] + dd[v]
		if total < ans {
			ans = total
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}
```
