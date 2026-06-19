package main

// LeetCode #1210: Minimum Moves to Reach Target with Rotations
// https://leetcode.com/problems/minimum-moves-to-reach-target-with-rotations/
// Difficulty: Hard
//
// A snake of length 2 (occupying 2 adjacent cells) moves through an n x n grid.
// The snake can move right, move down, rotate clockwise (horizontal->vertical),
// or rotate counter-clockwise (vertical->horizontal). Find the minimum number
// of moves to reach the target position: tail at (n-1, n-2) and head at
// (n-1, n-1), i.e., horizontal at the bottom-right corner.

import "fmt"

func main() {
	// Example 1
	grid := [][]int{
		{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0},
	}
	fmt.Println(minimumMoves(grid)) // 11

	// Example 2 (already at target)
	grid2 := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	fmt.Println(minimumMoves(grid2)) // 0

	// Blocked path
	grid3 := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(minimumMoves(grid3))

	// Impossible
	grid4 := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(minimumMoves(grid4))
}

// State represents the snake's configuration.
// (r, c) is the tail cell. dir=0 means horizontal (head at (r,c+1)),
// dir=1 means vertical (head at (r+1,c)).
type State struct {
	r, c, dir int
}

func minimumMoves(grid [][]int) int {
	n := len(grid)
	target := State{n - 1, n - 2, 0} // tail at bottom-left, horizontal

	// BFS
	dist := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = [2]int{-1, -1}
		}
	}

	queue := make([]State, 0)
	start := State{0, 0, 0}
	dist[0][0][0] = 0
	queue = append(queue, start)

	// Helper to check if a cell is free
	free := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0
	}

	directions := []struct {
		dr, dc, ndir int
		check        func(int, int) bool
	}{
		// dir == 0: horizontal (tail r,c, head r,c+1)
		// dir == 1: vertical (tail r,c, head r+1,c)
	}
	_ = directions

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		d := dist[cur.r][cur.c][cur.dir]

		if cur == target {
			return d
		}

		// Generate next states
		if cur.dir == 0 {
			// Horizontal
			// 1. Move right: tail goes to (r, c+1)
			if free(cur.r, cur.c+2) {
				ns := State{cur.r, cur.c + 1, 0}
				if dist[ns.r][ns.c][ns.dir] == -1 {
					dist[ns.r][ns.c][ns.dir] = d + 1
					queue = append(queue, ns)
				}
			}
			// 2. Move down: need both cells clear below
			if free(cur.r+1, cur.c) && free(cur.r+1, cur.c+1) {
				ns := State{cur.r + 1, cur.c, 0}
				if dist[ns.r][ns.c][ns.dir] == -1 {
					dist[ns.r][ns.c][ns.dir] = d + 1
					queue = append(queue, ns)
				}
			}
			// 3. Rotate clockwise: tail stays at (r,c), head goes to (r+1,c)
			//    Need (r+1,c) and (r+1,c+1) to be free (the 2x2 area)
			if free(cur.r+1, cur.c) && free(cur.r+1, cur.c+1) {
				ns := State{cur.r, cur.c, 1}
				if dist[ns.r][ns.c][ns.dir] == -1 {
					dist[ns.r][ns.c][ns.dir] = d + 1
					queue = append(queue, ns)
				}
			}
		} else {
			// Vertical
			// 1. Move right: need both cells to the right clear
			if free(cur.r, cur.c+1) && free(cur.r+1, cur.c+1) {
				ns := State{cur.r, cur.c + 1, 1}
				if dist[ns.r][ns.c][ns.dir] == -1 {
					dist[ns.r][ns.c][ns.dir] = d + 1
					queue = append(queue, ns)
				}
			}
			// 2. Move down: tail goes to (r+1, c)
			if free(cur.r+2, cur.c) {
				ns := State{cur.r + 1, cur.c, 1}
				if dist[ns.r][ns.c][ns.dir] == -1 {
					dist[ns.r][ns.c][ns.dir] = d + 1
					queue = append(queue, ns)
				}
			}
			// 3. Rotate counter-clockwise: tail stays at (r,c), head goes to (r,c+1)
			//    Need (r,c+1) and (r+1,c+1) to be free
			if free(cur.r, cur.c+1) && free(cur.r+1, cur.c+1) {
				ns := State{cur.r, cur.c, 0}
				if dist[ns.r][ns.c][ns.dir] == -1 {
					dist[ns.r][ns.c][ns.dir] = d + 1
					queue = append(queue, ns)
				}
			}
		}
	}

	return -1
}
