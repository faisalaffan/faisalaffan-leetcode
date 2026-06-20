# 2846 — Minimum Edge Weight Equilibrium Queries In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func minEdgeWeightEquilibriumQueries(n int, edges [][]int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2846: Minimum Edge Weight Equilibrium Queries in a Tree
// https://leetcode.com/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/
// Difficulty: Hard
//
// For each query (a,b), find minimum operations to make all edge weights on
// the path a->b equal. One operation changes any edge weight to any value.
// Equivalent to: total_edges_on_path - max_frequency_of_any_weight_on_path.
// Approach: root at 0, binary lifting for LCA, prefix frequency arrays for
// each weight (1-26) from root to each node.
// O(N * 26 + Q * (log N + 26)) time, O(N * 26 + N log N) space.

import "fmt"

func minEdgeWeightEquilibriumQueries(n int, edges [][]int, queries [][]int) []int {
	// Build adjacency list: neighbor, weight (0-indexed)
  // Matriks 2D
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]-1 // weight 1-indexed in input
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// Binary lifting setup
	LOG := 0
	for (1 << LOG) <= n {
		LOG++
	}
  // Matriks 2D
	up := make([][]int, n)
  // Alokasi slice
	depth := make([]int, n)
  // Alokasi slice
	freq := make([][26]int, n) // prefix freq from root to node

  // Range loop
	for i := range up {
		up[i] = make([]int, LOG)
	}

	// DFS from root 0
	var dfs func(u, p int)
	dfs = func(u, p int) {
		up[u][0] = p
		for j := 1; j < LOG; j++ {
			up[u][j] = up[up[u][j-1]][j-1]
		}
		for _, nei := range adj[u] {
			v, w := nei[0], nei[1]
			if v == p {
				continue
			}
			depth[v] = depth[u] + 1
			copy(freq[v][:], freq[u][:])
			freq[v][w]++
			dfs(v, u)
		}
	}
	// Initialize parent of root as root itself
	up[0][0] = 0
	dfs(0, 0)

	// LCA function
	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		diff := depth[u] - depth[v]
		for i := 0; i < LOG; i++ {
			if diff&(1<<i) != 0 {
				u = up[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := LOG - 1; i >= 0; i-- {
			if up[u][i] != up[v][i] {
				u = up[u][i]
				v = up[v][i]
			}
		}
		return up[u][0]
	}

  // Alokasi slice
	result := make([]int, len(queries))
	for qi, q := range queries {
		u, v := q[0], q[1]
		l := lca(u, v)
		totalEdges := depth[u] + depth[v] - 2*depth[l]
		maxFreq := 0
		for c := 0; c < 26; c++ {
			fc := freq[u][c] + freq[v][c] - 2*freq[l][c]
			if fc > maxFreq {
				maxFreq = fc
			}
		}
		result[qi] = totalEdges - maxFreq
	}
	return result
}

func main() {
	// Example 1
	n1 := 7
	edges1 := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 4, 2}, {4, 5, 2}, {5, 6, 2}}
	queries1 := [][]int{{0, 3}, {3, 6}, {2, 6}, {0, 6}}
	res1 := minEdgeWeightEquilibriumQueries(n1, edges1, queries1)
	for _, v := range res1 {
		fmt.Println(v)
	}

	// Example 2
	n2 := 8
	edges2 := [][]int{{1, 2, 6}, {1, 3, 4}, {2, 4, 6}, {2, 5, 3}, {3, 6, 6}, {3, 0, 8}, {7, 0, 2}}
	queries2 := [][]int{{4, 6}, {0, 4}, {6, 5}, {7, 4}}
	res2 := minEdgeWeightEquilibriumQueries(n2, edges2, queries2)
	for _, v := range res2 {
		fmt.Println(v)
	}

	// Single node edge case
	fmt.Println(minEdgeWeightEquilibriumQueries(1, [][]int{}, [][]int{{0, 0}}))

	// Two nodes
	fmt.Println(minEdgeWeightEquilibriumQueries(2, [][]int{{0, 1, 5}}, [][]int{{0, 1}}))

	// Query same node
	fmt.Println(minEdgeWeightEquilibriumQueries(3, [][]int{{0, 1, 1}, {1, 2, 2}}, [][]int{{0, 0}, {1, 1}}))
}
```
