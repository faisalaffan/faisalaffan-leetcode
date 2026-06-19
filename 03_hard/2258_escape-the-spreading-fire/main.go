package main

// LeetCode #2258: Escape the Spreading Fire
// https://leetcode.com/problems/escape-the-spreading-fire/
// Difficulty: Hard
//
// You are given a 2D grid where:
//   0 = grass (can walk), 1 = fire, 2 = wall
// Fire spreads to adjacent cells (4-directional) every minute.
// You start at (0, 0) and want to reach safehouse at (m-1, n-1).
// You can stay in place. You can't be on the same cell as fire at the same time.
// Return the maximum number of minutes you can wait before moving,
// while still being able to reach the safehouse before or at the same time as fire.
// If impossible, return -1. If unlimited, return 10^9.

import (
	"fmt"
	"math"
)

// maximumMinutes returns max wait time.
func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// fireDist[i][j] = minute when fire reaches (i,j), or INF
	fireDist := make([][]int, m)
	for i := range fireDist {
		fireDist[i] = make([]int, n)
		for j := range fireDist[i] {
			fireDist[i][j] = math.MaxInt32
		}
	}

	// BFS from all fire sources
	type point struct{ x, y int }
	queue := make([]point, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fireDist[i][j] = 0
				queue = append(queue, point{i, j})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nx, ny := p.x+d[0], p.y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n &&
				grid[nx][ny] != 2 && fireDist[nx][ny] == math.MaxInt32 {
				fireDist[nx][ny] = fireDist[p.x][p.y] + 1
				queue = append(queue, point{nx, ny})
			}
		}
	}

	// Binary search: can we wait `wait` minutes?
	canEscape := func(wait int) bool {
		// BFS for person
		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		pq := make([]point, 0)
		pq = append(pq, point{0, 0})
		visited[0][0] = true
		t := wait // time when we start moving

		for len(pq) > 0 {
			size := len(pq)
			for k := 0; k < size; k++ {
				p := pq[k]

				// check if we are at safehouse
				if p.x == m-1 && p.y == n-1 {
					return true
				}

				// we can't be on fire cell at this time (unless it's safehouse at same time)
				if fireDist[p.x][p.y] <= t {
					continue
				}

				for _, d := range dirs {
					nx, ny := p.x+d[0], p.y+d[1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n &&
						!visited[nx][ny] && grid[nx][ny] != 2 {

						// at the safehouse, fire can arrive at the same time
						if nx == m-1 && ny == n-1 && fireDist[nx][ny] <= t+1 {
							// we reach safehouse at t+1, fire reaches at <= t+1
							// Both arrive at same time? Actually fireDist <= t+1,
							// So fire might arrive before or at same time.
							// If fire arrives at same time, we're fine at safehouse.
							// But if fire arrives earlier, not fine.
							// Actually the problem says: you can be at safehouse
							// at the same time as fire.
							if fireDist[nx][ny] <= t {
								continue // fire already there
							}
							// fire arrives at t+1 or later -> ok
							visited[nx][ny] = true
							pq = append(pq, point{nx, ny})
						} else if fireDist[nx][ny] > t+1 {
							// we arrive at t+1, fire arrives later
							visited[nx][ny] = true
							pq = append(pq, point{nx, ny})
						}
					}
				}
			}
			pq = pq[size:]
			t++
		}
		return false
	}

	// If can't escape even with 0 wait
	if !canEscape(0) {
		return -1
	}

	// check unlimited (fire never reaches safehouse after waiting large)
	if canEscape(1_000_000_000) {
		return 1_000_000_000
	}

	// binary search
	lo, hi := 0, 1_000_000_000
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canEscape(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	// Example 1
	grid1 := [][]int{
		{0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 2, 1, 0},
		{0, 1, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(maximumMinutes(grid1)) // Expected: 1

	// Example 2
	grid2 := [][]int{
		{0, 0, 0, 0},
		{0, 1, 2, 0},
		{0, 2, 0, 0},
	}
	fmt.Println(maximumMinutes(grid2)) // Expected: -1

	// Example 3
	grid3 := [][]int{
		{0, 0, 0},
		{2, 2, 0},
		{1, 2, 0},
	}
	fmt.Println(maximumMinutes(grid3)) // Expected: 1000000000
}
