# 2065 — Maximum Path Quality Of A Graph

## Deskripsi

**Soal:** [2065. Maximum Path Quality Of A Graph](https://leetcode.com/problems/maximum-path-quality-of-a-graph/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), Backtracking (pelacakan mundur), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func maximalPathQuality(values []int, edges [][]int, maxTime int) int`

> **Ide Kunci:** DFS + Backtracking

## Solusi Go

```go
package main

// LeetCode #2065: Maximum Path Quality of a Graph
// https://leetcode.com/problems/maximum-path-quality-of-a-graph/
// Difficulty: Hard
// Approach: DFS + Backtracking

import "fmt"

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	// Build adjacency list: each entry is (neighbor, time)
  // Membuat slice 2D untuk DP/tabel
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, t})
		adj[v] = append(adj[v], [2]int{u, t})
	}

  // Membuat slice untuk menyimpan hasil
	visited := make([]int, n) // count visits to detect re-visits
	ans := 0

	var dfs func(u, time, quality int)
	dfs = func(u, time, quality int) {
		// Quality earned when first visiting a node
		if visited[u] == 0 {
			quality += values[u]
		}
		visited[u]++

		// If we're back at node 0, update answer
		if u == 0 {
			if quality > ans {
				ans = quality
			}
		}

		// Explore neighbors
		for _, edge := range adj[u] {
			v, t := edge[0], edge[1]
			if time+t <= maxTime {
				dfs(v, time+t, quality)
			}
		}

		// Backtrack
		visited[u]--
	}

	dfs(0, 0, 0)
	return ans
}

func main() {
	fmt.Println("2065. Maximum Path Quality of a Graph")

	// Example 1
	values1 := []int{0, 32, 10, 43}
	edges1 := [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}
	maxTime1 := 49
	fmt.Printf("values=%v edges=%v maxTime=%d → %d (expected 75)\n",
		values1, edges1, maxTime1, maximalPathQuality(values1, edges1, maxTime1))

	// Example 2
	values2 := []int{5, 10, 15, 20}
	edges2 := [][]int{{0, 1, 10}, {1, 2, 10}, {0, 3, 10}}
	maxTime2 := 30
	fmt.Printf("values=%v edges=%v maxTime=%d → %d (expected 25)\n",
		values2, edges2, maxTime2, maximalPathQuality(values2, edges2, maxTime2))

	// Example 3
	values3 := []int{1, 2, 3, 4}
	edges3 := [][]int{{0, 1, 10}, {1, 2, 11}, {2, 3, 12}, {1, 3, 13}}
	maxTime3 := 50
	fmt.Printf("values=%v edges=%v maxTime=%d → %d (expected 7)\n",
		values3, edges3, maxTime3, maximalPathQuality(values3, edges3, maxTime3))
}
```
