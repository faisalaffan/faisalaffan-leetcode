# 2421 — Number Of Good Paths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func newUF(n int) *uf
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2421: Number of Good Paths
// https://leetcode.com/problems/number-of-good-paths/
// Difficulty: Hard
//
// Union-Find by value. Sort nodes by value, union adjacent nodes that have
// <= current value. For each value group, count how many nodes in each
// connected component have that value. A good path starts and ends at nodes
// of the same value, and all intermediate nodes have <= that value.
// Time O(N log N + M alpha(N)) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(numberOfGoodPaths([]int{1, 3, 2, 1, 3},
		[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}))
	// Example 2
	fmt.Println(numberOfGoodPaths([]int{1, 1, 2, 2, 3},
		[][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}))
	// Single node
	fmt.Println(numberOfGoodPaths([]int{1}, [][]int{}))
}

type uf struct {
	parent []int
	rank   []int
}

func newUF(n int) *uf {
  // Alokasi slice integer
	p := make([]int, n)
  // Alokasi slice integer
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &uf{p, r}
}

func (u *uf) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *uf) union(x, y int) {
	xr, yr := u.find(x), u.find(y)
	if xr == yr {
		return
	}
	if u.rank[xr] < u.rank[yr] {
		xr, yr = yr, xr
	}
	u.parent[yr] = xr
	if u.rank[xr] == u.rank[yr] {
		u.rank[xr]++
	}
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	if n == 1 {
		return 1
	}

	// Build adjacency list
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// Sort nodes by value
  // Alokasi slice integer
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i
	}
  // Custom sort dengan comparator
	sort.Slice(nodes, func(i, j int) bool {
		return vals[nodes[i]] < vals[nodes[j]]
	})

	u := newUF(n)

	// For each value group, count good paths
	ans := n // each single node is a good path
	i := 0
	for i < n {
		j := i
		for j < n && vals[nodes[j]] == vals[nodes[i]] {
			j++
		}
		// Union all adjacent nodes with this value
		for k := i; k < j; k++ {
			node := nodes[k]
			for _, nei := range adj[node] {
				if vals[nei] <= vals[node] {
					u.union(node, nei)
				}
			}
		}
		// Count nodes in each component for this value
  // Membuat map (HashMap) — pencarian O(1)
		compCount := make(map[int]int)
		for k := i; k < j; k++ {
			root := u.find(nodes[k])
			compCount[root]++
		}
		for _, c := range compCount {
			ans += c * (c - 1) / 2
		}
		i = j
	}

	return ans
}
```
