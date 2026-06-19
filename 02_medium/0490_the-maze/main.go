package main

// LeetCode #490: The Maze
// https://leetcode.com/problems/the-maze/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{4, 4}))
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{3, 2}))
}

func TheMaze(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{start}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == destination[0] && cur[1] == destination[1] {
			return true
		}

		for _, d := range dirs {
			r, c := cur[0], cur[1]
			// Roll until hitting a wall
			for r+d[0] >= 0 && r+d[0] < m && c+d[1] >= 0 && c+d[1] < n && maze[r+d[0]][c+d[1]] == 0 {
				r += d[0]
				c += d[1]
			}
			if !visited[r][c] {
				visited[r][c] = true
				queue = append(queue, []int{r, c})
			}
		}
	}

	return false
}
