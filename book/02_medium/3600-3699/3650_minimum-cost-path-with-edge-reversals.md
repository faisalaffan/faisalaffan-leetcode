# 3650 — Minimum Cost Path With Edge Reversals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCostPathWithEdgeReversals(n int, edges [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O((n+m) log n)  
**Kompleksitas Ruang:** O(n+m)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3650: Minimum Cost Path with Edge Reversals
// https://leetcode.com/problems/minimum-cost-path-with-edge-reversals/
// Difficulty: Medium
// Time: O((n+m) log n) | Space: O(n+m)

import (
	"container/heap"
	"fmt"
	"math"
)

type edge struct {
	to, weight int
}

type minHeap []struct{ node, dist int }

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(struct{ node, dist int })) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumCostPathWithEdgeReversals(n int, edges [][]int) int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, 2 * w})
	}

  // Alokasi slice integer
	dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0

	pq := &minHeap{}
  // Masukkan elemen ke priority queue
	heap.Push(pq, struct{ node, dist int }{0, 0})

	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(pq).(struct{ node, dist int })
		u, d := cur.node, cur.dist
		if d > dist[u] {
			continue
		}
		if u == n-1 {
			return d
		}
		for _, e := range adj[u] {
			nd := d + e.weight
			if nd < dist[e.to] {
				dist[e.to] = nd
  // Masukkan elemen ke priority queue
				heap.Push(pq, struct{ node, dist int }{e.to, nd})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumCostPathWithEdgeReversals(4, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 2}}))
	fmt.Println(minimumCostPathWithEdgeReversals(3, [][]int{{0, 1, 10}, {1, 2, 5}}))
	fmt.Println(minimumCostPathWithEdgeReversals(4, [][]int{{0, 1, 1}, {0, 2, 4}, {1, 2, 1}, {2, 3, 2}}))
}
```
