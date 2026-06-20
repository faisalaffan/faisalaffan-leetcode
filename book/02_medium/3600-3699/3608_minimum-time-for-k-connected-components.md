# 3608 — Minimum Time For K Connected Components

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func MinimumTimeForKConnectedComponents(n int, edges [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3608: Minimum Time for K Connected Components
// https://leetcode.com/problems/minimum-time-for-k-connected-components/
// Difficulty: Medium
// Complexity: O(n + m) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}
	k := 2
	fmt.Println("Test 1:", MinimumTimeForKConnectedComponents(n, edges, k))
	// Test case 2
	n2 := 5
	edges2 := [][]int{{0, 1}, {2, 3}}
	k2 := 3
	fmt.Println("Test 2:", MinimumTimeForKConnectedComponents(n2, edges2, k2))
	// Test case 3
	n3 := 3
	edges3 := [][]int{{0, 1}, {1, 2}, {0, 2}}
	k3 := 1
	fmt.Println("Test 3:", MinimumTimeForKConnectedComponents(n3, edges3, k3))
}

func MinimumTimeForKConnectedComponents(n int, edges [][]int, k int) int {
	// Find connected components count
  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
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
	if components >= k {
		return 0
	}
	return k - components
}
```
