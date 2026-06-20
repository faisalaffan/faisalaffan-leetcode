# 3108 — Minimum Cost Walk In Weighted Graph

## Deskripsi

**Soal:** [3108. Minimum Cost Walk In Weighted Graph](https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** O(n + m + q * alpha(n))  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func NewDSU(n int) *DSU`

## Solusi Go

```go
package main

// LeetCode #3108: Minimum Cost Walk in Weighted Graph
// https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/
// Difficulty: Hard
// Time: O(n + m + q * alpha(n)) | Space: O(n)
//
// For each query (u,v), answer the minimum possible bitwise AND of a walk from u to v.
// Since AND only decreases with more edges, the minimum AND in a component is the AND
// of ALL edges in that component.

import (
	"fmt"
)

type DSU struct {
	parent []int
	and    []int
}

func NewDSU(n int) *DSU {
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	and := make([]int, n)
	mask := (1 << 30) - 1 // all 1s in lower 30 bits (max weight < 2^30)
	for i := 0; i < n; i++ {
		parent[i] = i
		and[i] = mask
	}
	return &DSU{parent: parent, and: and}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y, w int) {
	rx, ry := d.Find(x), d.Find(y)
	if rx == ry {
		d.and[rx] &= w
		return
	}
	d.and[rx] = d.and[rx] & d.and[ry] & w
	d.parent[ry] = rx
}

func (d *DSU) GetAnd(x int) int {
	return d.and[d.Find(x)]
}

func minimumCostWalk(n int, edges [][]int, query [][]int) []int {
	dsu := NewDSU(n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		dsu.Union(u, v, w)
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(query))
	for i, q := range query {
		u, v := q[0], q[1]
		if u == v {
			ans[i] = 0
		} else if dsu.Find(u) != dsu.Find(v) {
			ans[i] = -1
		} else {
			ans[i] = dsu.GetAnd(u)
		}
	}
	return ans
}

func main() {
	// Test case 1
	n := 5
	edges := [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}
	query := [][]int{{0, 3}, {3, 4}}
	fmt.Println("Test 1:", minimumCostWalk(n, edges, query))
	// Expected: [1, -1]

	// Test case 2
	n = 4
	edges = [][]int{{0, 0, 5}, {1, 2, 3}, {2, 3, 5}}
	query = [][]int{{0, 0}, {1, 3}, {0, 1}}
	fmt.Println("Test 2:", minimumCostWalk(n, edges, query))
	// Expected: [0, 1, -1]

	// Test case 3
	n = 3
	edges = [][]int{{0, 1, 7}, {1, 2, 3}}
	query = [][]int{{0, 2}}
	fmt.Println("Test 3:", minimumCostWalk(n, edges, query))
	// Expected: 7 & 3 = 3

	// Test case 4: single node, self query
	n = 1
	edges = [][]int{}
	query = [][]int{{0, 0}}
	fmt.Println("Test 4:", minimumCostWalk(n, edges, query))
	// Expected: [0]
}
```
