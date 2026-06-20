# 2368 — Reachable Nodes With Restrictions

## Deskripsi

**Soal:** [2368. Reachable Nodes With Restrictions](https://leetcode.com/problems/reachable-nodes-with-restrictions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #2368: Reachable Nodes With Restrictions
// https://leetcode.com/problems/reachable-nodes-with-restrictions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// BFS/DFS from node 0, skip restricted nodes.

import "fmt"

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5})) // 4
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {1, 5}, {2, 6}}, []int{1})) // 3
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
  // Membuat map untuk pencarian O(1): key → value
	restrictedSet := make(map[int]bool, len(restricted))
	for _, r := range restricted {
		restrictedSet[r] = true
	}

  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = true
		count := 1
		for _, v := range graph[u] {
			if !visited[v] && !restrictedSet[v] {
				count += dfs(v)
			}
		}
		return count
	}
	return dfs(0)
}
```
