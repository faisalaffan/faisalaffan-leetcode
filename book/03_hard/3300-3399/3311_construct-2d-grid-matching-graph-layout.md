# 3311 — Construct 2D Grid Matching Graph Layout

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func constructGridLayout(n int, edges [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3311: Construct 2D Grid Matching Graph Layout
// https://leetcode.com/problems/construct-2d-grid-matching-graph-layout/
// Difficulty: Hard
//
// Given an undirected graph with n nodes and edges, reconstruct a 2D grid such
// that every node appears exactly once and two nodes are adjacent in the grid
// iff there is an edge between them.
//
// Approach: Use degree analysis to identify corners (deg=2), edges (deg=3),
// and inner nodes (deg=4). Build the first row, then fill remaining rows.

import "fmt"

func main() {
	// Example 1
	fmt.Println(constructGridLayout(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}))
	// Example 2
	fmt.Println(constructGridLayout(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {1, 4}, {4, 3}}))
	// Example 3: single row
	fmt.Println(constructGridLayout(3, [][]int{{0, 1}, {1, 2}}))
	// Example 4: single column
	fmt.Println(constructGridLayout(3, [][]int{{0, 1}, {2, 1}}))
}

func constructGridLayout(n int, edges [][]int) [][]int {
	// Build adjacency list
  // Matriks 2D
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	// Find node by degree
  // Alokasi slice
	degNode := make([]int, 5)
  // Range loop
	for i := range degNode {
		degNode[i] = -1
	}
	for i := 0; i < n; i++ {
		degNode[len(g[i])] = i
	}

	// Build first row
  // Alokasi slice
	firstRow := make([]int, 0)

	if degNode[1] != -1 {
		// Single row: start from degree-1 node
		firstRow = append(firstRow, degNode[1])
	} else if degNode[4] == -1 {
		// Two columns: find two degree-2 nodes adjacent to each other
		x := degNode[2]
		for _, y := range g[x] {
			if len(g[y]) == 2 {
				firstRow = append(firstRow, x, y)
				break
			}
		}
	} else {
		// Multi-row: start from a corner (degree-2)
		x := degNode[2]
		firstRow = append(firstRow, x)
		prev := x
		x = g[x][0]
		for len(g[x]) > 2 {
			firstRow = append(firstRow, x)
			for _, y := range g[x] {
				if y != prev && len(g[y]) < 4 {
					prev = x
					x = y
					break
				}
			}
		}
		firstRow = append(firstRow, x)
	}

	cols := len(firstRow)
	rows := n / cols
  // Matriks 2D
	grid := make([][]int, rows)
	visited := make([]bool, n)

	for j := 0; j < cols; j++ {
		grid[0] = firstRow
		visited[firstRow[j]] = true
	}

	for r := 1; r < rows; r++ {
		grid[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			cur := grid[r-1][c]
			for _, nb := range g[cur] {
				if !visited[nb] {
					grid[r][c] = nb
					visited[nb] = true
					break
				}
			}
		}
	}

	return grid
}
```
