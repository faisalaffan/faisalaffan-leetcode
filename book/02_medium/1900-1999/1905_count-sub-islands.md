# 1905 — Count Sub Islands

## Deskripsi

**Soal:** [1905. Count Sub Islands](https://leetcode.com/problems/count-sub-islands/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n), Space: O(m*n) worst-case recursion  
**Kompleksitas Ruang:** O(m*n) worst-case recursion

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1905: Count Sub Islands
// https://leetcode.com/problems/count-sub-islands/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	fmt.Println(CountSubIslands(grid1, grid2))

	grid1b := [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}
	grid2b := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}
	fmt.Println(CountSubIslands(grid1b, grid2b))
}

// Time: O(m*n), Space: O(m*n) worst-case recursion
func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	count := 0

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || grid2[i][j] == 0 {
			return true
		}
		grid2[i][j] = 0
		result := grid1[i][j] == 1
		result = dfs(i-1, j) && result
		result = dfs(i+1, j) && result
		result = dfs(i, j-1) && result
		result = dfs(i, j+1) && result
		return result
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && dfs(i, j) {
				count++
			}
		}
	}
	return count
}
```
