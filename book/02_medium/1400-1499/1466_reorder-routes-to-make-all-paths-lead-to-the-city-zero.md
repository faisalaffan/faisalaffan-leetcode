# 1466 — Reorder Routes To Make All Paths Lead To The City Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minReorder(n int, connections [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n) where n = number of nodes  |  **Ruang:** O(n) for adjacency list

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1466: Reorder Routes to Make All Paths Lead to the City Zero
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}})) // 3

	// Test case 2
	fmt.Println(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}})) // 2

	// Test case 3
	fmt.Println(minReorder(3, [][]int{{1, 0}, {2, 0}})) // 0
}

// Time: O(n) where n = number of nodes
// Space: O(n) for adjacency list
func minReorder(n int, connections [][]int) int {
	// Build adjacency with direction info
	// For each edge, store [neighbor, direction]
	// direction=1 means original direction is away from 0, needs reorder
	// direction=0 means original direction is towards 0, ok
  // Matriks 2D
	adj := make([][][2]int, n)
	for _, conn := range connections {
		adj[conn[0]] = append(adj[conn[0]], [2]int{conn[1], 1}) // outgoing
		adj[conn[1]] = append(adj[conn[1]], [2]int{conn[0], 0}) // incoming
	}

	visited := make([]bool, n)
	changes := 0

	var dfs func(int)
	dfs = func(city int) {
		visited[city] = true
		for _, neighbor := range adj[city] {
			nextCity, needsReorder := neighbor[0], neighbor[1]
			if !visited[nextCity] {
				if needsReorder == 1 {
					changes++
				}
				dfs(nextCity)
			}
		}
	}

	dfs(0)
	return changes
}
```
