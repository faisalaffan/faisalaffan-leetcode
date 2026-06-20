# 1559 — Detect Cycles In 2D Grid

## Deskripsi

**Soal:** [1559. Detect Cycles In 2D Grid](https://leetcode.com/problems/detect-cycles-in-2d-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(R*C), Space: O(R*C)  
**Kompleksitas Ruang:** O(R*C)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1559: Detect Cycles in 2D Grid
// https://leetcode.com/problems/detect-cycles-in-2d-grid/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}
	fmt.Println(ContainsCycle(grid1))

	grid2 := [][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}
	fmt.Println(ContainsCycle(grid2))

	grid3 := [][]byte{{'a', 'b'}, {'b', 'a'}}
	fmt.Println(ContainsCycle(grid3))
}

func ContainsCycle(grid [][]byte) bool {
	// Time: O(R*C), Space: O(R*C)
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}

	rows, cols := len(grid), len(grid[0])
  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(r, c, pr, pc int) bool
	dfs = func(r, c, pr, pc int) bool {
		visited[r][c] = true

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if nr == pr && nc == pc {
				continue
			}
			if grid[nr][nc] != grid[r][c] {
				continue
			}
			if visited[nr][nc] {
				return true // cycle found
			}
			if dfs(nr, nc, r, c) {
				return true
			}
		}

		return false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] {
				if dfs(r, c, -1, -1) {
					return true
				}
			}
		}
	}

	return false
}
```
