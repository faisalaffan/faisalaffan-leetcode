# 2603 — Collect Coins In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func collectCoins(coins []int, edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2603: Collect Coins in a Tree
// https://leetcode.com/problems/collect-coins-in-a-tree/
// Difficulty: Hard

import "fmt"

// collectCoins calculates minimum moves to collect all coins starting from node 0.
// A coin at node v can be collected when the player is at any node u
// within distance 2 of v.
//
// Approach: Two-phase topological pruning.
// Phase 1: Recursively remove leaf nodes that have no coins (degree 1, coins=0).
//          These nodes cannot help collect coins since they're far from any coin.
// Phase 2: Remove one more layer of leaves (all remaining leaves regardless of coins).
//          Coins at these leaves can be collected from their neighbor (distance 1).
// Remaining edges must be traversed twice (go and return) to collect all coins.
//
// Complexity: O(n) time, O(n) space
func collectCoins(coins []int, edges [][]int) int {
	n := len(coins)
	if n <= 1 {
		return 0
	}

  // Matriks 2D
	adj := make([][]int, n)
  // Alokasi slice
	degree := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		degree[u]++
		degree[v]++
	}

	removed := make([]bool, n)

	// Phase 1: Remove leaf nodes with no coins (topological pruning)
  // Alokasi slice
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 1 && coins[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		removed[u] = true
		for _, v := range adj[u] {
			if !removed[v] {
				degree[v]--
				if degree[v] == 1 && coins[v] == 0 {
					q = append(q, v)
				}
			}
		}
	}

	// Phase 2: Remove one more layer (all remaining leaf nodes)
	q = make([]int, 0)
	for i := 0; i < n; i++ {
		if !removed[i] && degree[i] == 1 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		removed[u] = true
		for _, v := range adj[u] {
			if !removed[v] {
				degree[v]--
				// Don't push further; only one layer
			}
		}
	}

	// Count remaining edges (edges between non-removed nodes)
	remainingEdges := 0
	for _, e := range edges {
		if !removed[e[0]] && !removed[e[1]] {
			remainingEdges++
		}
	}

	return 2 * remainingEdges
}

func main() {
	// Test cases
	fmt.Println("Test 1: coins=[1,0,0,0,0,1], edges=[[0,1],[1,2],[2,3],[3,4],[4,5]] ->",
		collectCoins([]int{1, 0, 0, 0, 0, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}))

	fmt.Println("Test 2: coins=[0,0,0,1,1,0,0,1], edges=[[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]] ->",
		collectCoins([]int{0, 0, 0, 1, 1, 0, 0, 1}, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {5, 6}, {5, 7}}))

	fmt.Println("Test 3: coins=[1,1], edges=[[0,1]] ->",
		collectCoins([]int{1, 1}, [][]int{{0, 1}}))

	fmt.Println("Test 4: coins=[1], edges=[] ->",
		collectCoins([]int{1}, [][]int{}))

	fmt.Println("Test 5: coins=[0,0,0], edges=[[0,1],[1,2]] ->",
		collectCoins([]int{0, 0, 0}, [][]int{{0, 1}, {1, 2}}))

	fmt.Println("Test 6: coins=[0,0,1,0,0], edges=[[0,1],[1,2],[2,3],[3,4]] ->",
		collectCoins([]int{0, 0, 1, 0, 0}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
}
```
