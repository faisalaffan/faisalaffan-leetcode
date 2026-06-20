# 0827 — Making A Large Island

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func largestIsland(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #827: Making A Large Island
// https://leetcode.com/problems/making-a-large-island/
// Difficulty: Hard
//
// You are given an n x n binary matrix grid. You can change at most one 0 to 1.
// Return the size of the largest island (connected 1s, 4-directionally) after
// this operation. Return the original largest island size if no 0 exists.
//
// Approach: DFS labeling with Union-Find
//   - DFS to label each island with a unique ID (starting from 2), record its
//     size in a map.
//   - For each 0 cell, check up to 4 adjacent unique island IDs, sum their
//     sizes, add 1 for the flipped cell, track the maximum.

import "fmt"

func main() {
	// Example: [[1,1],[1,0]] → flipping (1,1) gives island size 4
	fmt.Println(largestIsland([][]int{
		{1, 1},
		{1, 0},
	}))

	// All 1s: 4
	fmt.Println(largestIsland([][]int{
		{1, 1},
		{1, 1},
	}))

	// Single cell: 1
	fmt.Println(largestIsland([][]int{
		{0},
	}))
	fmt.Println(largestIsland([][]int{
		{1},
	}))

	// No 0s: entire grid
	fmt.Println(largestIsland([][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}))

	// Disconnected islands with a single 0 bridge
	fmt.Println(largestIsland([][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}))
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Label each island with a unique ID (2, 3, 4, ...)
	id := 2
  // HashMap: O(1) lookup
	size := make(map[int]int)

	var dfs func(i, j, id int) int
	dfs = func(i, j, id int) int {
		if i < 0 || i >= n || j < 0 || j >= n || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = id
		count := 1
		for _, d := range dirs {
			count += dfs(i+d[0], j+d[1], id)
		}
		return count
	}

	// Label all islands
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				s := dfs(i, j, id)
				size[id] = s
				id++
			}
		}
	}

	// If already all 1s, return n*n
	if len(size) == 1 && size[2] == n*n {
		return n * n
	}

	// Track the largest existing island size (if we choose not to flip any 0)
	maxSize := 0
	for _, s := range size {
		if s > maxSize {
			maxSize = s
		}
	}

	// Try flipping each 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
  // HashMap: O(1) lookup
				seen := make(map[int]bool)
				total := 1 // the flipped cell
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < n {
						adjID := grid[ni][nj]
						if adjID >= 2 && !seen[adjID] {
							seen[adjID] = true
							total += size[adjID]
						}
					}
				}
				if total > maxSize {
					maxSize = total
				}
			}
		}
	}

	return maxSize
}
```
