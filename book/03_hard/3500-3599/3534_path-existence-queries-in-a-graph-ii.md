# 3534 — Path Existence Queries In A Graph Ii

## Deskripsi

**Soal:** [3534. Path Existence Queries In A Graph Ii](https://leetcode.com/problems/path-existence-queries-in-a-graph-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Process queries offline. Sort edges by weight, sort queries by

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	edges := make([]edge, 0)
	for i := 0; i < n-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff < 0 {
			diff = -diff
		}
		edges = append(edges, edge{i, i + 1, diff})
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// Process queries offline
	type query struct {
		idx int
		u, v int
	}
  // Membuat slice untuk menyimpan hasil
	qList := make([]query, len(queries))
	for i, q := range queries {
		qList[i] = query{i, q[0], q[1]}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))

	// Union-Find
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, n)
  // Iterasi seluruh elemen
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
