# 2642 — Design Graph With Shortest Path Calculator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int, edges [][]int) Graph
```

> **💡 Hint:** Floyd-Warshall for short paths.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Floyd-Warshall

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Floyd-Warshall** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2642: Design Graph With Shortest Path Calculator
// https://leetcode.com/problems/design-graph-with-shortest-path-calculator/
// Difficulty: Hard
//
// Approach: Floyd-Warshall for short paths.
// Since the graph is small (<= 100 nodes), Floyd-Warshall at construction
// and on each addEdge works well. On addEdge, re-run Floyd with the new
// edge as intermediate to propagate improvements.

import (
	"fmt"
	"math"
)

func main() {
	// Example: Graph(4, [[0,2,5],[0,1,2],[1,2,1],[3,0,3]])
	g := Constructor(4, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}})
	fmt.Println(g.ShortestPath(3, 2)) // 6
	fmt.Println(g.ShortestPath(0, 3)) // -1
	g.AddEdge([]int{1, 3, 4})
	fmt.Println(g.ShortestPath(0, 3)) // 6
}

type Graph struct {
	n    int
	dist [][]int
}

func Constructor(n int, edges [][]int) Graph {
	INF := math.MaxInt32
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w < dist[u][v] {
			dist[u][v] = w
		}
	}

	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					nd := dist[i][k] + dist[k][j]
					if nd < dist[i][j] {
						dist[i][j] = nd
					}
				}
			}
		}
	}

	return Graph{n: n, dist: dist}
}

func (g *Graph) AddEdge(edge []int) {
	u, v, w := edge[0], edge[1], edge[2]
	if w >= g.dist[u][v] {
		return
	}
	g.dist[u][v] = w

	// Re-run Floyd using only new edge as intermediate
	// For all (i,j), check if i->u->v->j is shorter
	INF := math.MaxInt32
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.dist[i][u] != INF && g.dist[v][j] != INF {
				nd := g.dist[i][u] + w + g.dist[v][j]
				if nd < g.dist[i][j] {
					g.dist[i][j] = nd
				}
			}
		}
	}
}

func (g *Graph) ShortestPath(node1 int, node2 int) int {
	if g.dist[node1][node2] == math.MaxInt32 {
		return -1
	}
	return g.dist[node1][node2]
}
```
