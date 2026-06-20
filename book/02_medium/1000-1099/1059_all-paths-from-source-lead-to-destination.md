# 1059 — All Paths From Source Lead To Destination

## Deskripsi

**Soal:** [1059. All Paths From Source Lead To Destination](https://leetcode.com/problems/all-paths-from-source-lead-to-destination/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V + E)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DFS with cycle detection. Every path from source must end at destination.

## Solusi Go

```go
package main

// LeetCode #1059: All Paths from Source Lead to Destination
// https://leetcode.com/problems/all-paths-from-source-lead-to-destination/
// Difficulty: Medium
//
// Approach: DFS with cycle detection. Every path from source must end at destination.
// Time: O(V + E)
// Space: O(V + E)

import "fmt"

func main() {
	fmt.Println(leadsToDestination(3, [][]int{{0, 1}, {0, 2}}, 0, 2)) // false
	fmt.Println(leadsToDestination(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 0, 3)) // true
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

  // Membuat slice untuk menyimpan hasil
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=processed

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] == 1 {
			return false // cycle
		}
		if state[node] == 2 {
			return true
		}

		if len(graph[node]) == 0 {
			return node == destination
		}

		state[node] = 1
		for _, next := range graph[node] {
			if !dfs(next) {
				return false
			}
		}
		state[node] = 2
		return true
	}

	return dfs(source)
}
```
