# 2646 — Minimize The Total Price Of The Trips

## Deskripsi

**Soal:** [2646. Minimize The Total Price Of The Trips](https://leetcode.com/problems/minimize-the-total-price-of-the-trips/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), DFS (Depth-First Search / pencarian kedalaman), Tree DP (DP pada pohon)

> **Ide Kunci:** Tree DP.

## Solusi Go

```go
package main

// LeetCode #2646: Minimize the Total Price of the Trips
// https://leetcode.com/problems/minimize-the-total-price-of-the-trips/
// Difficulty: Hard
//
// Approach: Tree DP.
// 1. For each trip, find the path and count how many times each node is visited.
// 2. Tree DP dp[node][0/1] = min total price for subtree rooted at node,
//    where 0 = node price NOT halved, 1 = node price IS halved.
//    Adjacent nodes cannot both be halved.

import "fmt"

func main() {
	// Example 1: n=4, edges=[[0,1],[1,2],[1,3]], price=[2,2,10,6], trips=[[0,3],[2,1],[2,3]] -> 23
	fmt.Println(minimumTotalPrice(4, [][]int{{0, 1}, {1, 2}, {1, 3}}, []int{2, 2, 10, 6}, [][]int{{0, 3}, {2, 1}, {2, 3}}))

	// Example 2: n=2, edges=[[0,1]], price=[2,2], trips=[[0,0]] -> 1
	fmt.Println(minimumTotalPrice(2, [][]int{{0, 1}}, []int{2, 2}, [][]int{{0, 0}}))
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Count visits per node
  // Membuat slice untuk menyimpan hasil
	cnt := make([]int, n)

	for _, trip := range trips {
		start, end := trip[0], trip[1]

		// DFS to find path from start to end
		var path []int
		var dfs func(u, parent int) bool
		dfs = func(u, parent int) bool {
			if u == end {
				path = append(path, u)
				return true
			}
			for _, v := range adj[u] {
				if v != parent {
					if dfs(v, u) {
						path = append(path, u)
						return true
					}
				}
			}
			return false
		}
		dfs(start, -1)

		for _, node := range path {
			cnt[node]++
		}
	}

	// Tree DP
	var dfs2 func(u, parent int) (int, int)
	dfs2 = func(u, parent int) (int, int) {
		notHalved := price[u] * cnt[u]
		halved := price[u] * cnt[u] / 2

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			childNot, childHalved := dfs2(v, u)
			// If u is halved, child cannot be halved
			halved += childNot
			// If u is not halved, child can be either (take min)
			notHalved += min(childNot, childHalved)
		}

		return notHalved, halved
	}

	a, b := dfs2(0, -1)
	return min(a, b)
}
```
