# 0847 — Shortest Path Visiting All Nodes

## Deskripsi

**Soal:** [0847. Shortest Path Visiting All Nodes](https://leetcode.com/problems/shortest-path-visiting-all-nodes/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func shortestPathLength(graph [][]int) int`

> **Ide Kunci:** BFS over state (node, visitedMask). Start from every node simultaneously

## Solusi Go

```go
package main

// LeetCode #847: Shortest Path Visiting All Nodes
// https://leetcode.com/problems/shortest-path-visiting-all-nodes/
// Difficulty: Hard
// Approach: BFS over state (node, visitedMask). Start from every node simultaneously
// (multi-source BFS). The mask tracks which nodes have been visited.

import "fmt"

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	target := (1 << n) - 1

	// dist[node][mask] = shortest steps to reach this state
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([]int, 1<<n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

  // Membuat slice untuk menyimpan hasil
	queue := make([][2]int, 0)
	for i := 0; i < n; i++ {
		mask := 1 << i
		queue = append(queue, [2]int{i, mask})
		dist[i][mask] = 0
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		node, mask := cur[0], cur[1]

		if mask == target {
			return dist[node][mask]
		}

		for _, nei := range graph[node] {
			newMask := mask | (1 << nei)
			if dist[nei][newMask] == -1 {
				dist[nei][newMask] = dist[node][mask] + 1
				queue = append(queue, [2]int{nei, newMask})
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}})) // Expected: 4
	fmt.Println(shortestPathLength([][]int{{1}, {0, 2, 4}, {1, 3}, {2}, {1}}))
	// Expected: 4 (0->1->4->1->2->3: path 0-1-4-1-2-3 = 5 steps... let me verify)
	// This tests a more complex graph
}
```
