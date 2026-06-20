# 1168 — Optimize Water Distribution In A Village

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewUnionFind(n int) *UnionFind
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1168: Optimize Water Distribution in a Village
// https://leetcode.com/problems/optimize-water-distribution-in-a-village/
// Difficulty: Hard (Premium)

import (
	"fmt"
	"sort"
)

// UnionFind for Kruskal's MST
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
  // Alokasi slice integer
	parent := make([]int, n)
  // Alokasi slice integer
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx == ry {
		return false
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
	return true
}

type Edge struct {
	u, v, w int
}

// minCostToSupplyWater uses a virtual node 0 connected to each house via well cost.
// Runs Kruskal's MST on n+1 nodes.
func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	edges := make([]Edge, 0, n+len(pipes))

	// Virtual node 0 to each house (1-indexed) = well cost
	for i, cost := range wells {
		edges = append(edges, Edge{0, i + 1, cost})
	}

	// Existing pipes
	for _, p := range pipes {
		edges = append(edges, Edge{p[0], p[1], p[2]})
	}

  // Custom sort dengan comparator
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w })

	uf := NewUnionFind(n + 1)
	total := 0
	for _, e := range edges {
		if uf.Union(e.u, e.v) {
			total += e.w
		}
	}
	return total
}

func main() {
	// Test case 1
	n := 3
	wells := []int{1, 2, 2}
	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}
	fmt.Println(minCostToSupplyWater(n, wells, pipes)) // 3

	// Test case 2: all wells cheaper than pipes
	n2 := 2
	wells2 := []int{1, 1}
	pipes2 := [][]int{{1, 2, 5}}
	fmt.Println(minCostToSupplyWater(n2, wells2, pipes2)) // 2

	// Test case 3: single house
	n3 := 1
	wells3 := []int{5}
	pipes3 := [][]int{}
	fmt.Println(minCostToSupplyWater(n3, wells3, pipes3)) // 5
}
```
