# 1489 — Find Critical And Pseudo Critical Edges In Minimum Spanning Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func newDSU(n int) *dSU`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Union-Find

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1489: Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
// https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
// Difficulty: Hard
//
// Approach: Kruskal per Edge
// 1. Compute MST weight normally.
// 2. For each edge i:
//    a. Forced-in: include edge i first, then run Kruskal. If weight > MST or
//       not all nodes connected → edge i is critical.
//    b. Forced-out: run Kruskal excluding edge i. If weight > MST or not all
//       connected → edge i is pseudo-critical (since excluding it makes MST worse,
//       but it's not critical if forced-in also yields MST weight).
// Return [critical, pseudoCritical].

import (
	"fmt"
	"sort"
)

func main() {
	n := 5
	edges := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 2},
		{0, 3, 2},
		{0, 4, 3},
		{3, 4, 3},
		{1, 4, 6},
	}
	result := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(result)
	// Expected: [[0,1],[2,3,4,5]]
}

type edge struct {
	u, v, w, idx int
}

type dSU struct {
	parent []int
	rank   []int
}

func newDSU(n int) *dSU {
  // Alokasi slice
	p := make([]int, n)
  // Alokasi slice
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &dSU{parent: p, rank: r}
}

func (d *dSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *dSU) union(x, y int) bool {
	x, y = d.find(x), d.find(y)
	if x == y {
		return false
	}
	if d.rank[x] < d.rank[y] {
		x, y = y, x
	}
	d.parent[y] = x
	if d.rank[x] == d.rank[y] {
		d.rank[x]++
	}
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	elist := make([]edge, m)
	for i, e := range edges {
		elist[i] = edge{u: e[0], v: e[1], w: e[2], idx: i}
	}
  // Custom sort
	sort.Slice(elist, func(i, j int) bool {
		return elist[i].w < elist[j].w
	})

	// Helper: compute MST weight with optional forced-in and forced-out edges.
	// forcedIn == nil means no forced edge.
	// forcedOut == -1 means no exclusion.
	mstWeight := func(forcedIn *edge, forcedOut int) int {
		dsu := newDSU(n)
		weight := 0
		edgesUsed := 0

		if forcedIn != nil {
			if dsu.union(forcedIn.u, forcedIn.v) {
				weight += forcedIn.w
				edgesUsed++
			}
		}

		for _, e := range elist {
			if e.idx == forcedOut {
				continue
			}
			if forcedIn != nil && e.idx == forcedIn.idx {
				continue
			}
			if dsu.union(e.u, e.v) {
				weight += e.w
				edgesUsed++
			}
		}

		if edgesUsed != n-1 {
			return -1 // not connected
		}
		return weight
	}

	baseWeight := mstWeight(nil, -1)

  // Alokasi slice
	critical := make([]int, 0)
  // Alokasi slice
	pseudo := make([]int, 0)

	for _, e := range elist {
		// Forced out (exclude this edge)
		wWithout := mstWeight(nil, e.idx)

		// If MST weight increases or graph disconnects, this edge is critical
		if wWithout == -1 || wWithout > baseWeight {
			critical = append(critical, e.idx)
			continue
		}

		// Forced in (include this edge)
		wWith := mstWeight(&e, -1)

		// If including it still gives MST weight, it's pseudo-critical
		// (it's in some MST but not all)
		if wWith == baseWeight {
			pseudo = append(pseudo, e.idx)
		}
	}

	return [][]int{critical, pseudo}
}
```
