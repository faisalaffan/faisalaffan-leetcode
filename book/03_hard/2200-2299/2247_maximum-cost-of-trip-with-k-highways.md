# 2247 — Maximum Cost Of Trip With K Highways

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func maximumCost(n int, highways [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2247: Maximum Cost of Trip With K Highways
// https://leetcode.com/problems/maximum-cost-of-trip-with-k-highways/
// Difficulty: Hard [Paid]
//
// Given n cities (0..n-1) connected by highways with toll costs.
// Find the maximum possible total cost of a trip that uses exactly k highways,
// visits each city at most once, and can start and end at any city.
// k <= n-1.

import (
	"fmt"
	"math"
)

// maximumCost returns maximum cost of a trip using exactly k highways.
func maximumCost(n int, highways [][]int, k int) int {
	if k >= n {
		return -1 // not enough cities for k highways (path of k edges needs k+1 cities)
	}
	if k == 0 {
		return 0 // no highways => cost 0
	}

	// build adjacency matrix (costs)
  // Matriks 2D
	adj := make([][][2]int, n) // [neighbor, cost]
	for _, h := range highways {
		u, v, cost := h[0], h[1], h[2]
		adj[u] = append(adj[u], [2]int{v, cost})
		adj[v] = append(adj[v], [2]int{u, cost})
	}

	// DP[mask][last] = max cost to reach 'last' city using mask of visited cities
	// mask has k+1 bits set (we visit k+1 cities for k highways)
  // Matriks 2D
	dp := make([][]int, 1<<n)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}

	// base: starting at each city (mask with 1 bit)
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 0
	}

	maxCost := math.MinInt32

	// iterate over all masks
	for mask := 1; mask < (1 << n); mask++ {
		bits := popcount(mask)

		if bits > k+1 {
			continue // too many cities
		}

		for last := 0; last < n; last++ {
			if dp[mask][last] == math.MinInt32 {
				continue
			}

			if bits == k+1 {
				// we have used exactly k edges (k+1 cities)
				if dp[mask][last] > maxCost {
					maxCost = dp[mask][last]
				}
				continue
			}

			// try extending
			for _, edge := range adj[last] {
				next, cost := edge[0], edge[1]
				if mask&(1<<next) != 0 {
					continue // already visited
				}
				newMask := mask | (1 << next)
				newCost := dp[mask][last] + cost
				if newCost > dp[newMask][next] {
					dp[newMask][next] = newCost
				}
			}
		}
	}

	if maxCost == math.MinInt32 {
		return -1
	}
	return maxCost
}

func popcount(x int) int {
	cnt := 0
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}

func main() {
	// Example 1
	n1 := 5
	highways1 := [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 7}, {3, 4, 2}, {0, 3, 18}}
	k1 := 3
	fmt.Println(maximumCost(n1, highways1, k1)) // Expected: 28

	// Example 2
	n2 := 4
	highways2 := [][]int{{0, 1, 3}, {2, 3, 2}}
	k2 := 2
	fmt.Println(maximumCost(n2, highways2, k2)) // Expected: -1
}
```
