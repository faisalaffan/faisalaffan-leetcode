# 2658 — Maximum Number Of Fish In A Grid

## Deskripsi

**Soal:** [2658. Maximum Number Of Fish In A Grid](https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func findMaxFish(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2658: Maximum Number of Fish in a Grid
// https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxFish := 0

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			return 0
		}
		fish := grid[r][c]
		grid[r][c] = 0 // Mark visited
		fish += dfs(r-1, c)
		fish += dfs(r+1, c)
		fish += dfs(r, c-1)
		fish += dfs(r, c+1)
		return fish
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				fish := dfs(i, j)
				if fish > maxFish {
					maxFish = fish
				}
			}
		}
	}
	return maxFish
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMaxFish([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", findMaxFish([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", findMaxFish([][]int{{0, 0}, {0, 0}}))
	// Expected: 0
}
```
