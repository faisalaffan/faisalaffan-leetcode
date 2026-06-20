# 1697 — Checking Existence Of Edge Length Limited Paths

## Deskripsi

**Soal:** [1697. Checking Existence Of Edge Length Limited Paths](https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func NewDSU(n int) *DSU`

## Solusi Go

```go
package main

// LeetCode #1697: Checking Existence of Edge Length Limited Paths
// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/
// Difficulty: Hard
// Strategy: Sort queries by limit, sort edges by weight.
// Union-Find to add edges incrementally as limits increase.

import (
	"fmt"
	"sort"
)

// Union-Find / DSU
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSU{parent, rank}
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

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// Sort edges by weight
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	// Sort queries by limit, keeping original indices
  // Membuat slice untuk menyimpan hasil
	q := make([][4]int, len(queries)) // [limit, u, v, originalIdx]
	for i, query := range queries {
		q[i] = [4]int{query[2], query[0], query[1], i}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][0] < q[j][0]
	})

	dsu := NewDSU(n)
  // Membuat slice untuk menyimpan hasil
	ans := make([]bool, len(queries))
	edgeIdx := 0

	for _, query := range q {
		limit, u, v, idx := query[0], query[1], query[2], query[3]
		// Add all edges with weight < limit
		for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < limit {
			dsu.Union(edgeList[edgeIdx][0], edgeList[edgeIdx][1])
			edgeIdx++
		}
		ans[idx] = dsu.Find(u) == dsu.Find(v)
	}
	return ans
}

func main() {
	// Example 1: n=3, edgeList=[[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries=[[0,1,2],[0,2,5]] -> [false,true]
	// (edge weight must be strictly less than limit)
	n1 := 3
	edgeList1 := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries1 := [][]int{{0, 1, 2}, {0, 2, 5}}
	fmt.Printf("distanceLimitedPathsExist(%d, %v, %v) = %v (expected [false true])\n",
		n1, edgeList1, queries1, distanceLimitedPathsExist(n1, edgeList1, queries1))

	// Example 2: n=5, edgeList=[[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries=[[0,4,14],[1,4,13]] -> [true,false]
	n2 := 5
	edgeList2 := [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}
	queries2 := [][]int{{0, 4, 14}, {1, 4, 13}}
	fmt.Printf("distanceLimitedPathsExist(%d, %v, %v) = %v (expected [true false])\n",
		n2, edgeList2, queries2, distanceLimitedPathsExist(n2, edgeList2, queries2))

	// Edge case: single query
	n3 := 2
	edgeList3 := [][]int{{0, 1, 5}}
	queries3 := [][]int{{0, 1, 3}}
	fmt.Printf("distanceLimitedPathsExist(%d, %v, %v) = %v (expected [false])\n",
		n3, edgeList3, queries3, distanceLimitedPathsExist(n3, edgeList3, queries3))
}
```
