package main

// LeetCode #1368: Minimum Cost to Make at Least One Valid Path in a Grid
// https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/
// Difficulty: Hard

import (
	"container/list"
	"fmt"
)

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 1:right, 2:left, 3:down, 4:up

	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 30
		}
	}
	dist[0][0] = 0

	// 0-1 BFS using deque
	dq := list.New()
	dq.PushBack([2]int{0, 0})

	for dq.Len() > 0 {
		front := dq.Remove(dq.Front()).([2]int)
		r, c := front[0], front[1]

		for dirIdx, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			cost := 0
			if grid[r][c] != dirIdx+1 {
				cost = 1
			}

			if dist[r][c]+cost < dist[nr][nc] {
				dist[nr][nc] = dist[r][c] + cost
				if cost == 0 {
					dq.PushFront([2]int{nr, nc})
				} else {
					dq.PushBack([2]int{nr, nc})
				}
			}
		}
	}

	return dist[m-1][n-1]
}

func main() {
	// Example 1
	fmt.Println(minCost([][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}))
	// Expected: 3

	// Example 2
	fmt.Println(minCost([][]int{
		{1, 1, 3},
		{3, 2, 2},
		{1, 1, 4},
	}))
	// Expected: 0

	// Example 3
	fmt.Println(minCost([][]int{
		{1, 2},
		{4, 3},
	}))
	// Expected: 1
}
