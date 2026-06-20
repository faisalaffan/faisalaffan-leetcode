# 3905 — Multi Source Flood Fill

## Deskripsi

**Soal:** [3905. Multi Source Flood Fill](https://leetcode.com/problems/multi-source-flood-fill/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N*M)  
**Kompleksitas Ruang:** O(N*M)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func MultiSourceFloodFill(n int, m int, sources [][]int) [][]int`

> **Ide Kunci:** Multi-source BFS. Track time and color for each cell.

## Solusi Go

```go
package main

// LeetCode #3905: Multi Source Flood Fill
// https://leetcode.com/problems/multi-source-flood-fill/
// Difficulty: Medium
// Time: O(N*M) | Space: O(N*M)
// Approach: Multi-source BFS. Track time and color for each cell.
// Use max color when multiple colors reach same cell at same time.

import "fmt"

func MultiSourceFloodFill(n int, m int, sources [][]int) [][]int {
  // Membuat slice 2D untuk DP/tabel
	grid := make([][]int, n)
  // Membuat slice 2D untuk DP/tabel
	time := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		time[i] = make([]int, m)
		for j := 0; j < m; j++ {
			time[i][j] = -1
		}
	}

	type Cell struct{ r, c, t, color int }
  // Membuat slice untuk menyimpan hasil
	queue := make([]Cell, 0)
	for _, src := range sources {
		r, c, color := src[0], src[1], src[2]
		grid[r][c] = color
		time[r][c] = 1
		queue = append(queue, Cell{r, c, 1, color})
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	head := 0
	for head < len(queue) {
		cur := queue[head]
		head++
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nr >= n || nc < 0 || nc >= m {
				continue
			}
			nt := cur.t + 1
			if time[nr][nc] == -1 {
				// Unvisited
				time[nr][nc] = nt
				grid[nr][nc] = cur.color
				queue = append(queue, Cell{nr, nc, nt, cur.color})
			} else if time[nr][nc] == nt && grid[nr][nc] < cur.color {
				// Same time, pick max color
				grid[nr][nc] = cur.color
			}
		}
	}

	return grid
}

func main() {
	// Example 1
	fmt.Println(MultiSourceFloodFill(3, 3, [][]int{{0, 0, 1}, {2, 2, 2}}))
	// Expected: [[1 1 2] [1 2 2] [2 2 2]]

	// Example 2
	fmt.Println(MultiSourceFloodFill(3, 3, [][]int{{0, 1, 3}, {1, 1, 5}}))
	// Expected: [[3 3 3] [5 5 5] [5 5 5]]

	// Example 3
	fmt.Println(MultiSourceFloodFill(2, 2, [][]int{{1, 1, 5}}))
	// Expected: [[5 5] [5 5]]
}
```
