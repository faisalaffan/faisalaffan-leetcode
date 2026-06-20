# 3772 — Maximum Subgraph Score In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func maxSubgraphScore(n int, edges [][]int, good []int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3772: Maximum Subgraph Score in a Tree
// https://leetcode.com/problems/maximum-subgraph-score-in-a-tree/
// Difficulty: Hard
//
// For each node, find max score of any connected subgraph containing it.
// Score = count(good nodes) - count(bad nodes) in the subgraph.
//
// Approach: Tree DP. First pass computes best downward score from each
// node. Second pass (reroot) computes best score for subgraphs containing
// each node by considering contributions from all neighbors.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxSubgraphScore(3, [][]int{{0, 1}, {0, 2}}, []int{1, 0, 1}))
	// Example 2
	fmt.Println(maxSubgraphScore(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, []int{1, 1, 0, 0}))
	// Edge: single node
	fmt.Println(maxSubgraphScore(1, [][]int{}, []int{1}))
	// Edge: all good
	fmt.Println(maxSubgraphScore(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, []int{1, 1, 1, 1}))
}

func maxSubgraphScore(n int, edges [][]int, good []int) []int64 {
	// Build adjacency
  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// good[i] = 1 (good) or 0 (bad)
  // Alokasi slice
	val := make([]int64, n)
  // Range loop
	for i := range val {
		if good[i] == 1 {
			val[i] = 1
		} else {
			val[i] = -1
		}
	}

	// dpDown[i] = best score of connected subgraph containing i,
	// restricted to i's subtree (downward only)
  // Alokasi slice
	dpDown := make([]int64, n)
	// ans[i] = final answer for node i
  // Alokasi slice
	ans := make([]int64, n)

	var dfs1 func(u, parent int)
	dfs1 = func(u, parent int) {
		dpDown[u] = val[u]
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			dfs1(v, u)
			if dpDown[v] > 0 {
				dpDown[u] += dpDown[v]
			}
		}
	}
	dfs1(0, -1)

	var dfs2 func(u, parent int, upScore int64)
	dfs2 = func(u, parent int, upScore int64) {
		ans[u] = val[u]
		if upScore > 0 {
			ans[u] += upScore
		}
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			if dpDown[v] > 0 {
				ans[u] += dpDown[v]
			}
		}

		// Compute upScore for children
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			// Score available above v = val[u] + (upScore if >0) +
			// sum of positive child contributions except v
			childUp := val[u]
			if upScore > 0 {
				childUp += upScore
			}
			for _, w := range adj[u] {
				if w == parent || w == v {
					continue
				}
				if dpDown[w] > 0 {
					childUp += dpDown[w]
				}
			}
			dfs2(v, u, childUp)
		}
	}
	dfs2(0, -1, 0)

	return ans
}
```
