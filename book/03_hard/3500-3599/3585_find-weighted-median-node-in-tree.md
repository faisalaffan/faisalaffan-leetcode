# 3585 — Find Weighted Median Node In Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func findWeightedMedianNode(n int, edges [][]int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3585: Find Weighted Median Node in Tree
// https://leetcode.com/problems/find-weighted-median-node-in-tree/
// Difficulty: Hard
//
// Find the node whose removal splits the tree into components
// each with weight <= totalWeight/2. Classic centroid with weights.
//
// Approach: Two DFS passes. First computes subtree sums. Second
// finds the node where max component weight <= total/2.

import "fmt"

func main() {
	// Example 1
	fmt.Println(findWeightedMedianNode(4, [][]int{{0, 1, 3}, {1, 2, 2}, {1, 3, 4}}, [][]int{{0, 1}}))
	// Example 2
	fmt.Println(findWeightedMedianNode(3, [][]int{{0, 1, 1}, {1, 2, 1}}, [][]int{{0, 2}}))
	// Edge: single node
	fmt.Println(findWeightedMedianNode(1, [][]int{}, [][]int{}))
}

func findWeightedMedianNode(n int, edges [][]int, queries [][]int) []int {
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}

  // Matriks 2D
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// First DFS for subtree sums from root 0
  // Alokasi slice
	subSum := make([]int64, n)
	total := int64(0)

	var dfs1 func(u, p int)
	dfs1 = func(u, p int) {
		for _, edge := range adj[u] {
			v, w := edge[0], edge[1]
			if v == p {
				continue
			}
			dfs1(v, u)
			subSum[u] += subSum[v] + int64(w)
		}
		if p == -1 {
			total = subSum[u]
		}
	}
	dfs1(0, -1)

	// Second DFS: find median node
  // Alokasi slice
	ans := make([]int, len(queries))
	for qi := 0; qi < len(queries); qi++ {
		// For this query, use current graph state
		// Simple: find centroid of current tree
		target := total / 2

		var centroid int
		var dfs2 func(u, p int)
		dfs2 = func(u, p int) {
			maxComp := total - subSum[u]
			for _, edge := range adj[u] {
				v, w := edge[0], edge[1]
				if v == p {
					continue
				}
				childSum := subSum[v] + int64(w)
				if childSum > maxComp {
					maxComp = childSum
				}
				dfs2(v, u)
			}
			if maxComp <= target {
				centroid = u
			}
		}
		dfs2(0, -1)
		ans[qi] = centroid
	}

	return ans
}
```
