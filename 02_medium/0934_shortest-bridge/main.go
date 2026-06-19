package main

// LeetCode #934: Shortest Bridge
// https://leetcode.com/problems/shortest-bridge/
// Difficulty: Medium

import "fmt"

// Time: O(n^2) | Space: O(n^2)
func shortestBridge(grid [][]int) int {
	n := len(grid)
	dirs := []int{1, 0, -1, 0, 1}
	q := make([][3]int, 0)

	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = 2
		for d := 0; d < 4; d++ {
			nx, ny := x+dirs[d], y+dirs[d+1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if grid[nx][ny] == 1 {
					dfs(nx, ny)
				} else if grid[nx][ny] == 0 {
					grid[nx][ny] = 2
					q = append(q, [3]int{nx, ny, 1})
				}
			}
		}
	}

	// DFS to mark first island
	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 2
				dfs(i, j)
				found = true
			}
		}
	}

	// BFS to find shortest path to second island
	for len(q) > 0 {
		cell := q[0]
		q = q[1:]
		for d := 0; d < 4; d++ {
			nx, ny := cell[0]+dirs[d], cell[1]+dirs[d+1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if grid[nx][ny] == 1 {
					return cell[2]
				}
				if grid[nx][ny] == 0 {
					grid[nx][ny] = 2
					q = append(q, [3]int{nx, ny, cell[2] + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(shortestBridge([][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}))
}
