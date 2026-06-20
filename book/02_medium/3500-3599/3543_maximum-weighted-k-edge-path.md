# 3543 — Maximum Weighted K Edge Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func MaximumWeightedKEdgePath(n int, edges [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3543: Maximum Weighted K-Edge Path
// https://leetcode.com/problems/maximum-weighted-k-edge-path/
// Difficulty: Medium
// Complexity: O(n + m*k) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 2}}
	k := 2
	fmt.Println("Test 1:", MaximumWeightedKEdgePath(n, edges, k))
	// Test case 2
	n2 := 3
	edges2 := [][]int{{0, 1, 10}, {1, 2, 20}}
	k2 := 2
	fmt.Println("Test 2:", MaximumWeightedKEdgePath(n2, edges2, k2))
	// Test case 3
	n3 := 2
	edges3 := [][]int{{0, 1, 100}}
	k3 := 1
	fmt.Println("Test 3:", MaximumWeightedKEdgePath(n3, edges3, k3))
}

func MaximumWeightedKEdgePath(n int, edges [][]int, k int) int {
	// Build adjacency list
  // Matriks 2D
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// dp[step][node] = max weight to reach node with exactly 'step' edges
  // Matriks 2D
	dp := make([][]int, k+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1 << 60
		}
	}
	dp[0][0] = 0

	for step := 1; step <= k; step++ {
		for u := 0; u < n; u++ {
			for _, edge := range adj[u] {
				v, w := edge[0], edge[1]
				if dp[step-1][u] != -1<<60 && dp[step-1][u]+w > dp[step][v] {
					dp[step][v] = dp[step-1][u] + w
				}
			}
		}
	}

	maxWeight := -1 << 60
	for u := 0; u < n; u++ {
		if dp[k][u] > maxWeight {
			maxWeight = dp[k][u]
		}
	}
	if maxWeight == -1<<60 {
		return -1
	}
	return maxWeight
}
```
