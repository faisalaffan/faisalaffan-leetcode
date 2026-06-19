package main

import (
	"fmt"
	"sort"
)

// LeetCode #1135: Connecting Cities With Minimum Cost
// https://leetcode.com/problems/connecting-cities-with-minimum-cost/

// Minimum cost to connect all cities (Kruskal's MST algorithm).

// Solution uses Union-Find and Kruskal's minimum spanning tree algorithm:
// 1. Sort edges by cost ascending
// 2. Connect cities using union-find, accumulating cost
// 3. If all cities are connected (single component), return cost; else -1

// Time complexity: O(E log E) = O(E log V) for sorting edges
// Space complexity: O(V) for the union-find data structure

type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	return true
}

func (uf *unionFind) connected() bool {
	root := uf.find(0)
	for i := 1; i < len(uf.parent); i++ {
		if uf.find(i) != root {
			return false
		}
	}
	return true
}

func minimumCost(n int, connections [][]int) int {
	if n <= 1 {
		return 0
	}

	// Sort by cost ascending
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	uf := newUnionFind(n)
	totalCost := 0
	edgesUsed := 0

	for _, conn := range connections {
		u, v, cost := conn[0]-1, conn[1]-1, conn[2]
		if uf.union(u, v) {
			totalCost += cost
			edgesUsed++
			if edgesUsed == n-1 {
				return totalCost
			}
		}
	}

	return -1
}

func main() {
	// Test case 1: Example from LeetCode
	n := 3
	connections := [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}
	fmt.Printf("minimumCost(%v, %v) = %d (expected: 6)\n", n, connections, minimumCost(n, connections))

	// Test case 2: Not all cities connected
	n2 := 4
	connections2 := [][]int{{1, 2, 3}, {3, 4, 4}}
	fmt.Printf("minimumCost(%v, %v) = %d (expected: -1)\n", n2, connections2, minimumCost(n2, connections2))

	// Test case 3: Single city
	n3 := 1
	connections3 := [][]int{}
	fmt.Printf("minimumCost(%v, %v) = %d (expected: 0)\n", n3, connections3, minimumCost(n3, connections3))
}
