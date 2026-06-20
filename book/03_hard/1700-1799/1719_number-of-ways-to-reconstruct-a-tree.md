# 1719 — Number Of Ways To Reconstruct A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkWays(pairs [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1719: Number Of Ways To Reconstruct A Tree
// https://leetcode.com/problems/number-of-ways-to-reconstruct-a-tree/
// Difficulty: Hard
// Strategy: Sort nodes by degree descending, assign parents from processed neighbors.
// Validate ancestor relationships. Check multiplicity when deg equals parent deg.

import (
	"fmt"
	"sort"
)

func checkWays(pairs [][]int) int {
	// Build adjacency and degrees
  // Membuat map (HashMap) — pencarian O(1)
	adj := make(map[int]map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	deg := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	nodeSet := make(map[int]bool)

	for _, p := range pairs {
		u, v := p[0], p[1]
		nodeSet[u] = true
		nodeSet[v] = true
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

	n := len(nodeSet)

	// Build sorted node list
  // Alokasi slice integer
	nodes := make([]int, 0, n)
	for node := range nodeSet {
		nodes = append(nodes, node)
	}
  // Custom sort dengan comparator
	sort.Slice(nodes, func(i, j int) bool {
		if deg[nodes[i]] != deg[nodes[j]] {
			return deg[nodes[i]] > deg[nodes[j]]
		}
		return nodes[i] < nodes[j]
	})

	maxDeg := deg[nodes[0]]

	// If max degree < n-1: only simple paths are valid (deg <= 2 for all nodes)
	if maxDeg < n-1 {
		if maxDeg > 2 {
			return 0
		}
		leafCount := 0
		for _, v := range nodes {
			if deg[v] > 2 {
				return 0
			}
			if deg[v] == 1 {
				leafCount++
			}
		}
		if leafCount != 2 {
			return 0
		}
		// Path graph: there is exactly 1 way
		return 1
	}

	// Root = first node (highest degree). No deg=n-1 requirement.
  // Membuat map (HashMap) — pencarian O(1)
	parent := make(map[int]int)
	parent[nodes[0]] = -1

	// Processing order for tiebreaking
  // Membuat map (HashMap) — pencarian O(1)
	order := make(map[int]int)
	for i, v := range nodes {
		order[v] = i
	}

	// Assign parents
	for _, v := range nodes[1:] {
		bestP := -1
		bestDeg := -1
		for u := range adj[v] {
			if _, ok := parent[u]; ok { // u is processed
				if deg[u] >= deg[v] {
					if bestP == -1 || deg[u] < bestDeg || (deg[u] == bestDeg && order[u] > order[bestP]) {
						bestP = u
						bestDeg = deg[u]
					}
				}
			}
		}
		if bestP == -1 {
			return 0
		}
		parent[v] = bestP
	}

	// Validate ancestor relationships for all pairs
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

	// Check multiplicity: if any node has same degree as parent
	// AND has another same-degree neighbor != parent
	for v, p := range parent {
		if p == -1 {
			continue
		}
		if deg[v] == deg[p] {
			for u := range adj[v] {
				if u != p && deg[u] == deg[v] {
					return 2
				}
			}
		}
	}

	return 1
}

func main() {
	// Example 1: [[1,2],[2,3]] -> 1
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", [][]int{{1, 2}, {2, 3}}, checkWays([][]int{{1, 2}, {2, 3}}))

	// Example 2: [[1,2],[2,3],[1,3]] -> 2
	fmt.Printf("checkWays(%v) = %d (expected 2)\n", [][]int{{1, 2}, {2, 3}, {1, 3}}, checkWays([][]int{{1, 2}, {2, 3}, {1, 3}}))

	// Example 3: [[1,2],[2,3],[2,4],[1,5]] -> 0
	fmt.Printf("checkWays(%v) = %d (expected 0)\n", [][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}, checkWays([][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}))

	// Linear chain [[1,2],[2,3],[3,4]] -> 1
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", [][]int{{1, 2}, {2, 3}, {3, 4}}, checkWays([][]int{{1, 2}, {2, 3}, {3, 4}}))

	// Star [[1,2],[1,3],[1,4]] -> 1
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", [][]int{{1, 2}, {1, 3}, {1, 4}}, checkWays([][]int{{1, 2}, {1, 3}, {1, 4}}))
}
```
