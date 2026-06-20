# 1579 — Remove Max Number Of Edges To Keep Graph Fully Traversable

## Deskripsi

**Soal:** [1579. Remove Max Number Of Edges To Keep Graph Fully Traversable](https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1579: Remove Max Number of Edges to Keep Graph Fully Traversable
// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
// Difficulty: Hard
//
// Union-Find approach:
// - Process type-3 (both) edges first: they benefit both Alice and Bob.
// - Then process type-1 (Alice) and type-2 (Bob) edges separately.
// - Count how many edges are actually used. Total edges - used = answer.
// - If either Alice or Bob can't fully traverse, return -1.

import "fmt"

func main() {
	// Example: n=4, edges=[[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]] -> 2
	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 2, 3},
		{1, 1, 3},
		{1, 2, 4},
		{1, 1, 2},
		{2, 3, 4},
	}))

	// Additional tests
	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 2, 3},
		{1, 1, 4},
		{2, 1, 4},
	}))

	fmt.Println(maxNumEdgesToRemove(2, [][]int{
		{1, 1, 2},
		{2, 1, 2},
		{3, 1, 2},
	}))

	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 3, 4},
		{1, 1, 3},
		{2, 2, 4},
	}))

	fmt.Println(maxNumEdgesToRemove(5, [][]int{
		{1, 1, 2},
		{2, 2, 3},
	}))
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(n int) *unionFind {
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, n+1)
  // Membuat slice untuk menyimpan hasil
	rank := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	return &unionFind{parent, rank, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	rx, ry := uf.find(x), uf.find(y)
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
	uf.count--
	return true
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice := newUnionFind(n)
	bob := newUnionFind(n)
	used := 0

	// Process type-3 edges first
	for _, e := range edges {
		if e[0] == 3 {
			connectedA := alice.union(e[1], e[2])
			connectedB := bob.union(e[1], e[2])
			if connectedA || connectedB {
				used++
			}
		}
	}

	// Process type-1 (Alice only)
	for _, e := range edges {
		if e[0] == 1 {
			if alice.union(e[1], e[2]) {
				used++
			}
		}
	}

	// Process type-2 (Bob only)
	for _, e := range edges {
		if e[0] == 2 {
			if bob.union(e[1], e[2]) {
				used++
			}
		}
	}

	// Check if both are fully traversable
	if alice.count != 1 || bob.count != 1 {
		return -1
	}

	return len(edges) - used
}
```
