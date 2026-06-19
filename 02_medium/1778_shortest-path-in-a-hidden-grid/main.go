package main

// LeetCode #1778: Shortest Path in a Hidden Grid
// https://leetcode.com/problems/shortest-path-in-a-hidden-grid/
// Difficulty: Medium [Paid]
// This is an interactive problem. We implement the solving algorithm.
// Time: O(m * n), Space: O(m * n)

import "fmt"

// GridMaster is the interface provided by LeetCode for the hidden grid problem.
// Real implementation: GridMaster.canMove(dir), GridMaster.move(dir), GridMaster.isTarget()
type GridMaster interface {
	CanMove(dir byte) bool
	Move(dir byte) bool // returns true if moved successfully
	IsTarget() bool
}

func findShortestPath(master GridMaster) int {
	// Discover the grid via DFS
	grid := make(map[[2]int]int) // 0=unvisited, 1=empty, 2=target, -1=blocked
	targetPos := [2]int{-1, -1}

	dirs := []byte{'U', 'D', 'L', 'R'}
	dr := map[byte]int{'U': -1, 'D': 1, 'L': 0, 'R': 0}
	dc := map[byte]int{'U': 0, 'D': 0, 'L': -1, 'R': 1}
	rev := map[byte]byte{'U': 'D', 'D': 'U', 'L': 'R', 'R': 'L'}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if master.IsTarget() {
			targetPos = [2]int{r, c}
			grid[[2]int{r, c}] = 2
		}
		for _, d := range dirs {
			nr, nc := r+dr[d], c+dc[d]
			key := [2]int{nr, nc}
			if _, visited := grid[key]; !visited && master.CanMove(d) {
				master.Move(d)
				grid[key] = 1
				dfs(nr, nc)
				// Move back
				master.Move(rev[d])
			} else if !master.CanMove(d) {
				grid[key] = -1 // blocked (wall)
			}
		}
	}

	grid[[2]int{0, 0}] = 1
	dfs(0, 0)

	if targetPos == [2]int{-1, -1} {
		return -1
	}

	// BFS for shortest path
	queue := [][2]int{{0, 0}}
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[i][0], queue[i][1]
			if r == targetPos[0] && c == targetPos[1] {
				return steps
			}
			for _, d := range dirs {
				nr, nc := r+dr[d], c+dc[d]
				key := [2]int{nr, nc}
				if v, ok := grid[key]; ok && v != -1 && !visited[key] {
					visited[key] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
		steps++
	}
	return -1
}

func main() {
	fmt.Println("Interactive problem - test via LeetCode platform")
	fmt.Println("Implementation ready for GridMaster interface")
}
