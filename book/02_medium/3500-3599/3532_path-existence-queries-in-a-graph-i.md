# 3532 — Path Existence Queries In A Graph I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewDSU(n int) *DSU
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3532: Path Existence Queries in a Graph I
// https://leetcode.com/problems/path-existence-queries-in-a-graph-i/
// Difficulty: Medium
// Complexity: O(n + q*alpha(n)) time, O(n) space

import "fmt"

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
  // Alokasi slice integer
	p := make([]int, n)
  // Alokasi slice integer
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	xr, yr := d.Find(x), d.Find(y)
	if xr == yr {
		return
	}
	if d.rank[xr] < d.rank[yr] {
		d.parent[xr] = yr
	} else if d.rank[xr] > d.rank[yr] {
		d.parent[yr] = xr
	} else {
		d.parent[yr] = xr
		d.rank[xr]++
	}
}

func main() {
	// Test case 1
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	queries := [][]int{{0, 2}, {0, 3}, {1, 4}}
	fmt.Println("Test 1:", PathExistenceQueriesInAGraphI(n, edges, queries))
	// Test case 2
	n2 := 3
	edges2 := [][]int{{0, 1}}
	queries2 := [][]int{{0, 1}, {1, 2}}
	fmt.Println("Test 2:", PathExistenceQueriesInAGraphI(n2, edges2, queries2))
	// Test case 3
	n3 := 2
	edges3 := [][]int{}
	queries3 := [][]int{{0, 1}}
	fmt.Println("Test 3:", PathExistenceQueriesInAGraphI(n3, edges3, queries3))
}

func PathExistenceQueriesInAGraphI(n int, edges [][]int, queries [][]int) []bool {
	dsu := NewDSU(n)
	for _, e := range edges {
		dsu.Union(e[0], e[1])
	}
	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = dsu.Find(q[0]) == dsu.Find(q[1])
	}
	return result
}
```
