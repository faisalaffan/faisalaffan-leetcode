package main

// LeetCode #3603: Minimum Cost Path with Alternating Directions II
// https://leetcode.com/problems/minimum-cost-path-with-alternating-directions-ii/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Test case 1
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Test 1:", MinimumCostPathWithAlternatingDirectionsIi(grid))
	// Test case 2
	grid2 := [][]int{{1, 1}, {1, 1}}
	fmt.Println("Test 2:", MinimumCostPathWithAlternatingDirectionsIi(grid2))
	// Test case 3
	grid3 := [][]int{{5}}
	fmt.Println("Test 3:", MinimumCostPathWithAlternatingDirectionsIi(grid3))
}

func MinimumCostPathWithAlternatingDirectionsIi(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return grid[0][0]
	}

	// 0=even direction (horizontal), 1=odd direction (vertical), 2=start
	dist := make([][][3]int, m)
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

	dist[0][0][2] = grid[0][0]
	queue.PushBack(state{0, 0, 2})

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).(state)
		r, c, dir := front.r, front.c, front.dir

		// After horizontal (dir=0), next must be vertical (dir=1) and vice versa
		if dir != 0 {
			// Can move horizontally to adjacent cells
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
			// Can move vertically to adjacent cells
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
