# 0882 — Reachable Nodes In Subdivided Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func reachableNodes(edges [][]int, maxMoves int, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #882: Reachable Nodes In Subdivided Graph
// https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/
// Difficulty: Hard
//
// Dijkstra from node 0. Each original edge [u, v, cnt] represents cnt
// subdivided nodes between u and v. Distance weight is cnt+1 to traverse
// the full edge (cnt subdivided + 1 destination original).
//
// After Dijkstra, for each edge:
//   - Reachable from u: max(0, maxMoves - dist[u]) steps into the edge
//   - Reachable from v: max(0, maxMoves - dist[v]) steps into the edge
//   - Total subdivided nodes on this edge = min(cnt, from_u + from_v)
//
// +1 for each original node that is reachable (dist <= maxMoves).

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	dist, node int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	// Build adjacency list: edge weight = cnt + 1 (to traverse the full edge)
  // Membuat matriks/slice 2D untuk DP
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		w := cnt + 1
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// Dijkstra
  // Alokasi slice integer
	dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	pq := &MinHeap{{0, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(pq).(Item)
		d, u := cur.dist, cur.node
		if d > dist[u] {
			continue
		}
		for _, edge := range adj[u] {
			v, w := edge[0], edge[1]
			nd := d + w
			if nd < dist[v] {
				dist[v] = nd
  // Masukkan elemen ke priority queue
				heap.Push(pq, Item{nd, v})
			}
		}
	}

	// Count reachable original nodes.
	result := 0
	for i := 0; i < n; i++ {
		if dist[i] <= maxMoves {
			result++
		}
	}

	// Count reachable subdivided nodes on each edge.
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		fromU := max(0, maxMoves-dist[u])
		fromV := max(0, maxMoves-dist[v])
		result += min(cnt, fromU+fromV)
	}
	return result
}

func main() {
	// Example 1: edges=[[0,1,10],[0,2,1],[1,2,2]], maxMoves=6, n=3 -> 13
	edges1 := [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}
	fmt.Println("Test 1:", reachableNodes(edges1, 6, 3)) // 13

	// Example 2: edges=[[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves=10, n=4 -> 23
	edges2 := [][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}}
	fmt.Println("Test 2:", reachableNodes(edges2, 10, 4)) // 23

	// Edge: maxMoves=0, can only reach node 0
	edges3 := [][]int{{0, 1, 5}}
	fmt.Println("Test 3:", reachableNodes(edges3, 0, 2)) // 1 (only node 0)

	// Edge: actual subdivided nodes counted correctly
	edges4 := [][]int{{0, 1, 5}}
	fmt.Println("Test 4:", reachableNodes(edges4, 3, 2)) // ? Let's compute:
	// dist[0]=0, dist[1]=6 (cnt+1=6)
	// fromU=min(3, maxMoves-0)=3, fromV=0
	// result = 1 (node 0) + min(5, 3+0) = 1+3 = 4
	fmt.Println("Test 5:", reachableNodes(edges4, 6, 2)) // dist[1]=6 <=6, so both nodes + all 5 subdivided
	// 2 + 5 = 7
}
```
