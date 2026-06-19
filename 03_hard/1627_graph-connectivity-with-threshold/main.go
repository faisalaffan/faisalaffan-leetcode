package main

// LeetCode #1627: Graph Connectivity With Threshold
// https://leetcode.com/problems/graph-connectivity-with-threshold/
// Difficulty: Hard

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
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

func (uf *UnionFind) Union(x, y int) {
	xr, yr := uf.Find(x), uf.Find(y)
	if xr == yr {
		return
	}
	if uf.rank[xr] < uf.rank[yr] {
		uf.parent[xr] = yr
	} else if uf.rank[xr] > uf.rank[yr] {
		uf.parent[yr] = xr
	} else {
		uf.parent[yr] = xr
		uf.rank[xr]++
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func areConnected(n int, threshold int, queries [][]int) []bool {
	uf := NewUnionFind(n + 1) // 1-indexed

	// For each divisor d > threshold, connect all multiples of d
	for d := threshold + 1; d <= n; d++ {
		for m := 2 * d; m <= n; m += d {
			uf.Union(d, m)
		}
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = uf.Connected(q[0], q[1])
	}
	return result
}

func main() {
	// Test case 1: n=6, threshold=2, queries=[[1,4],[2,5],[3,6]] -> [false,false,true]
	n := 6
	threshold := 2
	queries := [][]int{{1, 4}, {2, 5}, {3, 6}}
	result := areConnected(n, threshold, queries)
	fmt.Printf("n=%d threshold=%d queries=%v -> %v (expected [false,false,true])\n",
		n, threshold, queries, result)

	// Test case 2: n=6, threshold=0, queries=[[4,5],[3,4],[3,2],[2,6],[1,3]] -> [true,false,true,false,true]
	n2 := 6
	threshold2 := 0
	queries2 := [][]int{{4, 5}, {3, 4}, {3, 2}, {2, 6}, {1, 3}}
	result2 := areConnected(n2, threshold2, queries2)
	fmt.Printf("n=%d threshold=%d queries=%v -> %v\n", n2, threshold2, queries2, result2)

	// Test case 3: n=5, threshold=1, queries=[[4,5],[4,5],[3,2],[2,3],[3,4]] -> [false,false,false,false,false]
	n3 := 5
	threshold3 := 1
	queries3 := [][]int{{4, 5}, {4, 5}, {3, 2}, {2, 3}, {3, 4}}
	result3 := areConnected(n3, threshold3, queries3)
	fmt.Printf("n=%d threshold=%d queries=%v -> %v\n", n3, threshold3, queries3, result3)
}
