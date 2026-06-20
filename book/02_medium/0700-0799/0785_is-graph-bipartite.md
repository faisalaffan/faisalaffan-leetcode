# 0785 — Is Graph Bipartite

## Deskripsi

**Soal:** [0785. Is Graph Bipartite](https://leetcode.com/problems/is-graph-bipartite/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #785: Is Graph Bipartite?
// https://leetcode.com/problems/is-graph-bipartite/
// Difficulty: Medium
// Time: O(V + E)
// Space: O(V)

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
  // Membuat slice untuk menyimpan hasil
	color := make([]int, n)

	var dfs func(node int, c int) bool
	dfs = func(node int, c int) bool {
		if color[node] != 0 {
			return color[node] == c
		}
		color[node] = c

		for _, neighbor := range graph[node] {
			if !dfs(neighbor, -c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
```
