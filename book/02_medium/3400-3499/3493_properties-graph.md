# 3493 — Properties Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func PropertiesGraph(properties [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3493: Properties Graph
// https://leetcode.com/problems/properties-graph/
// Difficulty: Medium
// Complexity: O(n * p + n^2) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	props := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println("Test 1:", PropertiesGraph(props, 1))

	// Test case 2
	props2 := [][]int{{1, 1}, {1, 1}, {1, 1}}
	fmt.Println("Test 2:", PropertiesGraph(props2, 2))

	// Test case 3
	props3 := [][]int{{1}, {2}, {3}}
	fmt.Println("Test 3:", PropertiesGraph(props3, 0))
}

func PropertiesGraph(properties [][]int, k int) int {
	n := len(properties)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Build adjacency: nodes share >= k common properties
  // Matriks 2D
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
  // HashMap: O(1) lookup
		setI := make(map[int]bool)
		for _, v := range properties[i] {
			setI[v] = true
		}
		for j := i + 1; j < n; j++ {
			common := 0
			for _, v := range properties[j] {
				if setI[v] {
					common++
				}
			}
			if common >= k {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	// DFS to count connected components
	visited := make([]bool, n)
	components := 0
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfs(i)
		}
	}
	return components
}
```
