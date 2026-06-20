# 2812 — Find The Safest Path In A Grid

## Deskripsi

**Soal:** [2812. Find The Safest Path In A Grid](https://leetcode.com/problems/find-the-safest-path-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2 log n)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** Binary Search (pencarian biner), BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func FindTheSafestPathInAGrid(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2812: Find the Safest Path in a Grid
// https://leetcode.com/problems/find-the-safest-path-in-a-grid/
// Difficulty: Medium
// Time: O(n^2 log n) | Space: O(n^2)

import (
	"fmt"
	"math"
)

func FindTheSafestPathInAGrid(grid [][]int) int {
	n := len(grid)

	// Multi-source BFS to compute distance to nearest thief (1)
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	type cell struct{ r, c int }
  // Membuat slice untuk menyimpan hasil
	queue := make([]cell, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				queue = append(queue, cell{i, j})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && dist[nr][nc] > dist[cur.r][cur.c]+1 {
				dist[nr][nc] = dist[cur.r][cur.c] + 1
				queue = append(queue, cell{nr, nc})
			}
		}
	}

	// Binary search for max safety factor
	canReach := func(minDist int) bool {
		if dist[0][0] < minDist {
			return false
		}
  // Membuat slice 2D untuk DP/tabel
		visited := make([][]bool, n)
  // Iterasi seluruh elemen
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		visited[0][0] = true
		q := []cell{{0, 0}}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur.r == n-1 && cur.c == n-1 {
				return true
			}
			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] && dist[nr][nc] >= minDist {
					visited[nr][nc] = true
					q = append(q, cell{nr, nc})
				}
			}
		}
		return false
	}

	lo, hi := 0, n*2
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canReach(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}

func main() {
	fmt.Println(FindTheSafestPathInAGrid([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))
	fmt.Println(FindTheSafestPathInAGrid([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}))
}
```
