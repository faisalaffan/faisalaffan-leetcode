package main

import (
	"fmt"
)

// LeetCode #1162: As Far from Land as Possible
// https://leetcode.com/problems/as-far-from-land-as-possible/
// Difficulty: Medium

// Multi-source BFS from all land cells simultaneously.
// The last water cell to be reached has the max distance.

// Time: O(m*n)
// Space: O(m*n)

func maxDistance(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return -1
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := make([][2]int, 0)

	// Add all land cells to queue
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// If all water or all land, return -1
	if len(queue) == 0 || len(queue) == n*n {
		return -1
	}

	distance := -1
	for len(queue) > 0 {
		distance++
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
					grid[nr][nc] = 1 // mark visited
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	return distance
}

func main() {
	grid1 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	fmt.Printf("%d (expected: 2)\n", maxDistance(grid1))

	grid2 := [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Printf("%d (expected: 4)\n", maxDistance(grid2))

	grid3 := [][]int{{0, 0}, {0, 0}}
	fmt.Printf("%d (expected: -1)\n", maxDistance(grid3))
}
