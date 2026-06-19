package main

// LeetCode #1719: Number Of Ways To Reconstruct A Tree
// https://leetcode.com/problems/number-of-ways-to-reconstruct-a-tree/
// Difficulty: Hard
// Strategy: Sort nodes by degree. Root has deg=n-1.
// For each node, its parent is the neighbor with smallest degree >= its own.
// Check validity. If any node has same degree as parent -> 2 ways.

import (
	"fmt"
	"sort"
)

func checkWays(pairs [][]int) int {
	// Build adjacency sets and degree counts
	adj := make(map[int]map[int]bool)
	deg := make(map[int]int)
	nodes := make(map[int]bool)

	for _, p := range pairs {
		u, v := p[0], p[1]
		nodes[u] = true
		nodes[v] = true
		if adj[u] == nil {
			adj[u] = make(map[int]bool)
		}
		if adj[v] == nil {
			adj[v] = make(map[int]bool)
		}
		adj[u][v] = true
		adj[v][u] = true
		deg[u]++
		deg[v]++
	}

	n := len(nodes)

	// Build node list and sort by degree descending
	nodeList := make([]int, 0, n)
	for node := range nodes {
		nodeList = append(nodeList, node)
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return deg[nodeList[i]] > deg[nodeList[j]]
	})

	// Root must have degree n-1
	if deg[nodeList[0]] != n-1 {
		return 0
	}

	// Map node -> parent
	parent := make(map[int]int)

	// For each node (except root), find its parent:
	// the neighbor with smallest degree >= its own degree
	for _, v := range nodeList {
		if deg[v] == n-1 {
			parent[v] = -1
			continue
		}
		// Find candidate parent
		bestParent := -1
		for u := range adj[v] {
			if deg[u] >= deg[v] {
				if bestParent == -1 || deg[u] < deg[bestParent] || (deg[u] == deg[bestParent] && u < bestParent) {
					bestParent = u
				}
			}
		}
		if bestParent == -1 {
			return 0
		}
		parent[v] = bestParent
	}

	// Verify: for every pair (u,v), the parent-child relationships are consistent
	// In the tree, for every edge in pairs, one must be ancestor of the other
	// Check parent chains
	isAncestor := func(anc, desc int) bool {
		for desc != -1 {
			if desc == anc {
				return true
			}
			desc = parent[desc]
		}
		return false
	}

	for _, p := range pairs {
		u, v := p[0], p[1]
		if !isAncestor(u, v) && !isAncestor(v, u) {
			return 0
		}
	}

	// Check for multiple ways: if any node has the same degree as its parent
	// (and parent is not root), there are 2 ways
	ways := 1
	for v, p := range parent {
		if p == -1 {
			continue
		}
		if deg[v] == deg[p] {
			ways = 2
			break
		}
	}

	return ways
}

func main() {
	// Example 1: [[1,2],[2,3]] -> 1
	pairs1 := [][]int{{1, 2}, {2, 3}}
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", pairs1, checkWays(pairs1))

	// Example 2: [[1,2],[2,3],[1,3]] -> 2
	pairs2 := [][]int{{1, 2}, {2, 3}, {1, 3}}
	fmt.Printf("checkWays(%v) = %d (expected 2)\n", pairs2, checkWays(pairs2))

	// Example 3: [[1,2],[2,3],[2,4],[1,5]] -> 0
	pairs3 := [][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}
	fmt.Printf("checkWays(%v) = %d (expected 0)\n", pairs3, checkWays(pairs3))

	// Linear chain [[1,2],[2,3],[3,4]] -> 1
	pairs4 := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", pairs4, checkWays(pairs4))

	// Star: [[1,2],[1,3],[1,4]] -> 1
	pairs5 := [][]int{{1, 2}, {1, 3}, {1, 4}}
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", pairs5, checkWays(pairs5))
}
