# 3596 — Minimum Cost Path With Alternating Directions I

## Deskripsi

**Soal:** [3596. Minimum Cost Path With Alternating Directions I](https://leetcode.com/problems/minimum-cost-path-with-alternating-directions-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dijkstra (lintasan terpendek), LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #3596: Minimum Cost Path with Alternating Directions I
// https://leetcode.com/problems/minimum-cost-path-with-alternating-directions-i/
// Difficulty: Medium [Paid]
// Complexity: O(n*m) time, O(n*m) space

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Test case 1
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 1:", MinimumCostPathWithAlternatingDirectionsI(grid))
	// Test case 2
	grid2 := [][]int{{1, 1, 1}, {1, 1, 1}}
	fmt.Println("Test 2:", MinimumCostPathWithAlternatingDirectionsI(grid2))
	// Test case 3
	grid3 := [][]int{{5}}
	fmt.Println("Test 3:", MinimumCostPathWithAlternatingDirectionsI(grid3))
}

func MinimumCostPathWithAlternatingDirectionsI(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return grid[0][0]
	}

	// Dijkstra-like with direction (0=horizontal, 1=vertical, 2=start)
  // Membuat slice 2D untuk DP/tabel
	dist := make([][][3]int, m)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([][3]int, n)
		for j := range dist[i] {
			for k := range dist[i][j] {
				dist[i][j][k] = math.MaxInt32
			}
		}
	}

	type state struct {
		r, c, dir int
	}
	queue := list.New()

	// Start
	dist[0][0][2] = grid[0][0]
	queue.PushBack(state{0, 0, 2})

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).(state)
		r, c, dir := front.r, front.c, front.dir

		if dir != 0 {
			// Can move horizontally
			for _, nc := range []int{c - 1, c + 1} {
				if nc >= 0 && nc < n {
					nd := dist[r][c][dir] + grid[r][nc]
					if nd < dist[r][nc][0] {
						dist[r][nc][0] = nd
						queue.PushBack(state{r, nc, 0})
					}
				}
			}
		}
		if dir != 1 {
			// Can move vertically
			for _, nr := range []int{r - 1, r + 1} {
				if nr >= 0 && nr < m {
					nd := dist[r][c][dir] + grid[nr][c]
					if nd < dist[nr][c][1] {
						dist[nr][c][1] = nd
						queue.PushBack(state{nr, c, 1})
					}
				}
			}
		}
	}

	result := dist[m-1][n-1][0]
	if dist[m-1][n-1][1] < result {
		result = dist[m-1][n-1][1]
	}
	if dist[m-1][n-1][2] < result {
		result = dist[m-1][n-1][2]
	}
	return result
}
```
