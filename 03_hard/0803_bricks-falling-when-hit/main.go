package main

// LeetCode #803: Bricks Falling When Hit
// https://leetcode.com/problems/bricks-falling-when-hit/
// Difficulty: Hard
//
// Given a grid of 1s (bricks) and 0s (empty), we fire a sequence of hits.
// A brick is stable if it is connected to the top row. Bricks that become
// disconnected after a hit fall (turn to 0).
//
// Approach: Reverse Union-Find
//   - Mark hit cells as "to be removed" (2). Build the final graph of remaining
//     bricks. Use Union-Find to mark all stable bricks (connected to a "top"
//     sentinel).
//   - Process hits in reverse: restore each hit brick, union with neighbors,
//     count how many new bricks became connected to the top. That count is the
//     number of bricks that would have fallen at this hit.

import "fmt"

func main() {
	// Example: grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
	// After hit at (1,0), brick falls → [1]
	fmt.Println(hitBricks([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 0},
	}, [][]int{{1, 0}}))

	// Example from problem:
	// grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]] → [0, 0]
	fmt.Println(hitBricks([][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
	}, [][]int{{1, 1}, {1, 0}}))

	// Example from problem: grid=[[1,0,1],[1,1,1]], hits=[[0,0],[0,2],[1,1]] → [0,3,0]
	fmt.Println(hitBricks([][]int{
		{1, 0, 1},
		{1, 1, 1},
	}, [][]int{{0, 0}, {0, 2}, {1, 1}}))

	// Example: grid = [[1,1,1],[0,1,0],[0,0,0]], hits = [[0,2],[1,1],[0,0]] → [0,0,1]
	fmt.Println(hitBricks([][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 0},
	}, [][]int{{0, 2}, {1, 1}, {0, 0}}))
}

func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])

	// Mark hit cells: 1 → 2 (to be removed)
	for _, hit := range hits {
		if grid[hit[0]][hit[1]] == 1 {
			grid[hit[0]][hit[1]] = 2
		}
	}

	// Union-Find with extra sentinel index for "top" (row 0)
	parent := make([]int, m*n+1)
	size := make([]int, m*n+1)
	for i := range parent {
		parent[i] = i
		if i < m*n {
			size[i] = 1
		}
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		rx, ry := find(x), find(y)
		if rx == ry {
			return
		}
		if size[rx] < size[ry] {
			rx, ry = ry, rx
		}
		parent[ry] = rx
		size[rx] += size[ry]
	}

	top := m * n
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Union all remaining bricks after hits
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				idx := i*n + j
				if i == 0 {
					union(idx, top)
				}
				if i > 0 && grid[i-1][j] == 1 {
					union(idx, (i-1)*n+j)
				}
				if j > 0 && grid[i][j-1] == 1 {
					union(idx, i*n+j-1)
				}
			}
		}
	}

	// Process hits in reverse
	result := make([]int, len(hits))
	for k := len(hits) - 1; k >= 0; k-- {
		i, j := hits[k][0], hits[k][1]
		if grid[i][j] != 2 {
			continue // was 0 originally — no brick to restore
		}

		grid[i][j] = 1 // restore
		idx := i*n + j

		before := size[find(top)]

		if i == 0 {
			union(idx, top)
		}
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
				union(idx, ni*n+nj)
			}
		}

		after := size[find(top)]
		added := after - before - 1 // subtract the brick itself
		if added < 0 {
			added = 0
		}
		result[k] = added
	}

	return result
}
