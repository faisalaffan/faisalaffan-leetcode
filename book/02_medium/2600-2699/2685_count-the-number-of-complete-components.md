# 2685 — Count The Number Of Complete Components

## Deskripsi

**Soal:** [2685. Count The Number Of Complete Components](https://leetcode.com/problems/count-the-number-of-complete-components/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V + E)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func countCompleteComponents(n int, edges [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2685: Count the Number of Complete Components
// https://leetcode.com/problems/count-the-number-of-complete-components/
// Difficulty: Medium
// Time: O(V + E) | Space: O(V + E)

import "fmt"

func countCompleteComponents(n int, edges [][]int) int {
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		// BFS/DFS to find component
		queue := []int{i}
		visited[i] = true
		vertices := []int{}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			vertices = append(vertices, u)
			for _, v := range adj[u] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}

		// Check if complete: each vertex should have (k-1) edges to other vertices in component
		k := len(vertices)
		isComplete := true
		for _, v := range vertices {
			if len(adj[v]) != k-1 {
				isComplete = false
				break
			}
		}
		if isComplete {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
	// Expected: 2 ({0,1,2} complete, {3,4} not complete with 5)

	// Test case 2
	fmt.Println("Test 2:", countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}))
	// Expected: 1

	// Test case 3: single nodes
	fmt.Println("Test 3:", countCompleteComponents(3, [][]int{}))
	// Expected: 3 (each single node is complete)
}
```
