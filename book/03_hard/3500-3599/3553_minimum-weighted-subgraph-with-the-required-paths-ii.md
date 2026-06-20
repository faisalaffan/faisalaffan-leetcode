# 3553 — Minimum Weighted Subgraph With The Required Paths Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func dijkstra(adj [][]Edge, start int) []int64
```

> **💡 Hint:** Dijkstra from src1, src2, and reverse Dijkstra from dest.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3553: Minimum Weighted Subgraph With the Required Paths II
// https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths-ii/
// Difficulty: Hard
//
// Given a weighted directed graph and three nodes src1, src2, dest, find
// the minimum weight of a subgraph that contains paths from src1 to dest
// and from src2 to dest.
//
// Approach: Dijkstra from src1, src2, and reverse Dijkstra from dest.
// For each node, total = dist1[node] + dist2[node] + distRev[node].

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minimumWeightedSubgraph(6, [][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}}, 0, 1, 5))
	// Example 2
	fmt.Println(minimumWeightedSubgraph(3, [][]int{{0, 1, 1}, {1, 2, 2}}, 0, 1, 2))
	// Edge: single node
	fmt.Println(minimumWeightedSubgraph(1, [][]int{}, 0, 0, 0))
}

type Edge struct {
	to, weight int
}
type Item struct {
	node    int
	dist    int64
}
type PriorityQueue []Item
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} { old := *pq; n := len(old); x := old[n-1]; *pq = old[:n-1]; return x }

func dijkstra(adj [][]Edge, start int) []int64 {
	n := len(adj)
  // Alokasi slice integer
	dist := make([]int64, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0
	pq := &PriorityQueue{}
  // Masukkan elemen ke priority queue
	heap.Push(pq, Item{start, 0})
	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		item := heap.Pop(pq).(Item)
		u := item.node
		if item.dist > dist[u] {
			continue
		}
		for _, e := range adj[u] {
			if nd := item.dist + int64(e.weight); nd < dist[e.to] {
				dist[e.to] = nd
  // Masukkan elemen ke priority queue
				heap.Push(pq, Item{e.to, nd})
			}
		}
	}
	return dist
}

func minimumWeightedSubgraph(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]Edge, n)
  // Membuat matriks/slice 2D untuk DP
	radj := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], Edge{v, w})
		radj[v] = append(radj[v], Edge{u, w})
	}

	dist1 := dijkstra(adj, src1)
	dist2 := dijkstra(adj, src2)
	distD := dijkstra(radj, dest)

	ans := int64(math.MaxInt64)
	for i := 0; i < n; i++ {
		if dist1[i] != math.MaxInt64 && dist2[i] != math.MaxInt64 && distD[i] != math.MaxInt64 {
			total := dist1[i] + dist2[i] + distD[i]
			if total < ans {
				ans = total
			}
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}
```
