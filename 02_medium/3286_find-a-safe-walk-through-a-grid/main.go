package main

// LeetCode #3286: Find a Safe Walk Through a Grid
// https://leetcode.com/problems/find-a-safe-walk-through-a-grid/
// Difficulty: Medium
// Time: O(m * n) Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(findSafeWalk([][]int{{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}, 1)) // true
	fmt.Println(findSafeWalk([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 1)) // true
	fmt.Println(findSafeWalk([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, 3)) // true
}

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// minHealthLost[i][j] = minimum health lost to reach (i,j)
	minLost := make([][]int, m)
	for i := range minLost {
		minLost[i] = make([]int, n)
		for j := range minLost[i] {
			minLost[i][j] = 1 << 30
		}
	}
	minLost[0][0] = grid[0][0]

	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur[0], cur[1]
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				cost := minLost[i][j] + grid[x][y]
				if cost < minLost[x][y] {
					minLost[x][y] = cost
					queue = append(queue, [2]int{x, y})
				}
			}
		}
	}

	return minLost[m-1][n-1] < health
}
