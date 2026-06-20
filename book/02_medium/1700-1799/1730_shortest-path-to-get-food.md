# 1730 — Shortest Path To Get Food

## Deskripsi

**Soal:** [1730. Shortest Path To Get Food](https://leetcode.com/problems/shortest-path-to-get-food/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n), Space: O(m * n)  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** —

**Fungsi Solusi:** `func getFood(grid [][]byte) int`

## Solusi Go

```go
package main

// LeetCode #1730: Shortest Path to Get Food
// https://leetcode.com/problems/shortest-path-to-get-food/
// Difficulty: Medium [Paid]
// Time: O(m * n), Space: O(m * n)

import "fmt"

func getFood(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	startR, startC := 0, 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '*' {
				startR, startC = r, c
			}
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
  // Membuat slice 2D untuk DP/tabel
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	queue := [][2]int{{startR, startC}}
	visited[startR][startC] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			if grid[r][c] == '#' {
				return steps
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] != 'X' {
					visited[nr][nc] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
		steps++
	}
	return -1
}

func main() {
	grid1 := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X'},
	}
	fmt.Println(getFood(grid1)) // Expected: 3

	grid2 := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'X', 'O', 'X'},
		{'X', 'O', 'X', '#', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
	}
	fmt.Println(getFood(grid2)) // Expected: -1

	grid3 := [][]byte{
		{'*', 'O', '#'},
	}
	fmt.Println(getFood(grid3)) // Expected: 2
}
```
