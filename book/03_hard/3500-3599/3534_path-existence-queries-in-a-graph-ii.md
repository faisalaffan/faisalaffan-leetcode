# 3534 — Path Existence Queries In A Graph Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int
```

> **💡 Hint:** Process queries offline. Sort edges by weight, sort queries by

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3534: Path Existence Queries in a Graph II
// https://leetcode.com/problems/path-existence-queries-in-a-graph-ii/
// Difficulty: Hard
//
// Given a graph, queries ask if path exists where max edge weight difference
// <= maxDiff. Process dynamic queries with DSU rollback or persistent DSU.
//
// Approach: Process queries offline. Sort edges by weight, sort queries by
// maxDiff, use union-find to connect edges as threshold increases.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(pathExistenceQueries(3, []int{1, 2, 3}, 2, [][]int{{0, 2}}))
	// Example 2
	fmt.Println(pathExistenceQueries(4, []int{1, 4, 2, 3}, 1, [][]int{{0, 1}, {2, 3}}))
	// Edge: single node
	fmt.Println(pathExistenceQueries(1, []int{5}, 10, [][]int{{0, 0}}))
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	// Build edges between adjacent nodes with weight diff
	type edge struct {
		u, v, w int
	}
	edges := make([]edge, 0)
	for i := 0; i < n-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff < 0 {
			diff = -diff
		}
		edges = append(edges, edge{i, i + 1, diff})
	}

	// Sort edges by weight
  // Custom sort dengan comparator
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// Process queries offline
	type query struct {
		idx int
		u, v int
	}
	qList := make([]query, len(queries))
	for i, q := range queries {
		qList[i] = query{i, q[0], q[1]}
	}

  // Alokasi slice integer
	ans := make([]int, len(queries))

	// Union-Find
  // Alokasi slice integer
	parent := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			parent[x] = y
		}
	}

	ei := 0
	// For each query (in sorted order of maxDiff)
	for _, q := range qList {
		u, v := q.u, q.v
		// Connect edges with weight <= maxDiff
		for ei < len(edges) && edges[ei].w <= maxDiff {
			union(edges[ei].u, edges[ei].v)
			ei++
		}
		if find(u) == find(v) {
			ans[q.idx] = 1
		} else {
			ans[q.idx] = 0
		}
	}

	return ans
}
```
