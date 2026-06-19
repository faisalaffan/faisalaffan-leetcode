package main

// LeetCode #286: Walls and Gates
// https://leetcode.com/problems/walls-and-gates/
// Difficulty: Medium [Paid]
// Time: O(m*n), Space: O(m*n) for queue

import "fmt"

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	rows, cols := len(rooms), len(rooms[0])
	queue := [][2]int{}
	INF := 2147483647

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			r, c := cell[0]+dir[0], cell[1]+dir[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && rooms[r][c] == INF {
				rooms[r][c] = rooms[cell[0]][cell[1]] + 1
				queue = append(queue, [2]int{r, c})
			}
		}
	}
}

func main() {
	INF := 2147483647
	rooms1 := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	wallsAndGates(rooms1)
	fmt.Println(rooms1)

	rooms2 := [][]int{{INF}}
	wallsAndGates(rooms2)
	fmt.Println(rooms2)
}
