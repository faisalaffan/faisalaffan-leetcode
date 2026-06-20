# 0802 — Find Eventual Safe States

## Deskripsi

**Soal:** [0802. Find Eventual Safe States](https://leetcode.com/problems/find-eventual-safe-states/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #802: Find Eventual Safe States
// https://leetcode.com/problems/find-eventual-safe-states/
// Difficulty: Medium
// Time: O(V + E)
// Space: O(V)

import "fmt"

func main() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
  // Membuat slice untuk menyimpan hasil
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=safe

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] > 0 {
			return state[node] == 2
		}

		state[node] = 1
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		state[node] = 2
		return true
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if dfs(i) {
			result = append(result, i)
		}
	}

	return result
}
```
