# 1568 — Minimum Number Of Days To Disconnect Island

## Deskripsi

**Soal:** [1568. Minimum Number Of Days To Disconnect Island](https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** // 1. Count islands. If not exactly 1, return 0 (already disconnected).

## Solusi Go

```go
package main

// LeetCode #1568: Minimum Number of Days to Disconnect Island
// https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/
// Difficulty: Hard
//
// Approach:
// 1. Count islands. If not exactly 1, return 0 (already disconnected).
// 2. Try flipping each land cell to water. If it disconnects the island, return 1.
// 3. Otherwise, return 2 (worst case: remove any corner of an island).
//
// Key insight: answer is always 0, 1, or 2. If removing one cell doesn't work,
// removing two cells always works (e.g., remove the two ends of a bridge).

import "fmt"

func main() {
	// Example: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]] -> 2
	fmt.Println(minDays([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}))

	// Example: single cell island -> 1
	fmt.Println(minDays([][]int{{1}}))

	// Example: already disconnected -> 0
	fmt.Println(minDays([][]int{
		{1, 1},
		{1, 0},
	}))

	// Example: 2x2 all land -> 2
	fmt.Println(minDays([][]int{
		{1, 1},
		{1, 1},
	}))

	// Example: line of 3 -> 1 (remove middle)
	fmt.Println(minDays([][]int{{1, 1, 1}}))
}

var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Count islands
	if countIslands(grid) != 1 {
		return 0
	}

	// Try removing each land cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if countIslands(grid) != 1 {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}

	return 2
}

func countIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, m)
  // Iterasi seluruh elemen
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count++
				dfs(grid, visited, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]int, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 || visited[i][j] {
		return
	}
	visited[i][j] = true
	for _, d := range dirs {
		dfs(grid, visited, i+d[0], j+d[1])
	}
}
```
