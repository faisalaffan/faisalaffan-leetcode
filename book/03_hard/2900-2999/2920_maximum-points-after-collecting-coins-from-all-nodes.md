# 2920 — Maximum Points After Collecting Coins From All Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumPoints(edges [][]int, coins []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2920: Maximum Points After Collecting Coins From All Nodes
// https://leetcode.com/problems/maximum-points-after-collecting-coins-from-all-nodes/
// Difficulty: Hard
//
// Tree DP with memoization. At each node, we can either:
//   1. Collect coins[i] - k points (pay penalty)
//   2. Halve coins[i] (floor division by 2) and collect the reduced value
// Halving at a node also halves coins in the entire subtree because the
// "shifts" count propagates downward. Since coins[i] <= 10^4, at most 14
// halvings reduce everything to 0.
// DP[node][shifts] = max points from subtree when coins have been halved
// `shifts` times before reaching this node.

import "fmt"

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	const maxShifts = 16 // enough for coins up to 10^4 (log2(10000) ≈ 14)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, maxShifts)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(node, parent, shifts int) int
	dfs = func(node, parent, shifts int) int {
		if shifts >= maxShifts {
			return 0
		}
		if dp[node][shifts] != -1 {
			return dp[node][shifts]
		}

		// Option 1: collect coins with penalty k at this node
		collect := (coins[node] >> shifts) - k
		// Option 2: halve coins at this node (collect halved value)
		halve := coins[node] >> (shifts + 1)

		for _, child := range g[node] {
			if child == parent {
				continue
			}
			collect += dfs(child, node, shifts)     // children not halved
			halve += dfs(child, node, shifts+1)       // children also halved
		}

		if collect > halve {
			dp[node][shifts] = collect
		} else {
			dp[node][shifts] = halve
		}
		return dp[node][shifts]
	}

	return dfs(0, -1, 0)
}

func main() {
	// Example: edges=[[0,1],[1,2],[2,3]], coins=[10,10,3,3], k=2
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}}, []int{10, 10, 3, 3}, 2))

	// Same tree, k=5
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}}, []int{10, 10, 3, 3}, 5))

	// Star tree, k=0
	fmt.Println(maximumPoints([][]int{{0, 1}, {0, 2}}, []int{8, 4, 4}, 0))

	// Single node
	fmt.Println(maximumPoints([][]int{}, []int{5}, 2))

	// All coins small
	fmt.Println(maximumPoints([][]int{{0, 1}}, []int{0, 0}, 1))

	// Larger k (penalty large, always better to halve)
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}}, []int{10, 10, 10}, 20))

	// k=0, always collect
	fmt.Println(maximumPoints([][]int{{0, 1}, {0, 2}}, []int{5, 3, 7}, 0))
}
```
