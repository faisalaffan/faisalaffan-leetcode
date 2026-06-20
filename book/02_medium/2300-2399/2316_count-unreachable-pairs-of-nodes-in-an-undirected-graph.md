# 2316 — Count Unreachable Pairs Of Nodes In An Undirected Graph

## Deskripsi

**Soal:** [2316. Count Unreachable Pairs Of Nodes In An Undirected Graph](https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func countPairs(n int, edges [][]int) int64`

## Solusi Go

```go
package main

// LeetCode #2316: Count Unreachable Pairs of Nodes in an Undirected Graph
// https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func countPairs(n int, edges [][]int) int64 {
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)
	var totalPairs int64 = 0
	var prevComponents int64 = 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		// BFS/DFS to find component size
		size := int64(0)
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			size++
			for _, nei := range graph[node] {
				if !visited[nei] {
					visited[nei] = true
					queue = append(queue, nei)
				}
			}
		}
		totalPairs += prevComponents * size
		prevComponents += size
	}
	return totalPairs
}

func main() {
	// Test case 1
	fmt.Println(countPairs(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	// Expected: 0

	// Test case 2
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
	// Expected: 14
}
```
