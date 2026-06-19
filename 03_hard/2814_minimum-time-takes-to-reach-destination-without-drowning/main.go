package main

// LeetCode #2814: Minimum Time Takes to Reach Destination Without Drowning
// https://leetcode.com/problems/minimum-time-takes-to-reach-destination-without-drowning/
// Difficulty: Hard [Paid]
//
// Grid with "S" start, "D" destination, "." empty, "X" stone, "*" water.
// Each second you can move 4-directionally. Simultaneously water spreads from
// each "*" to adjacent "." cells. You cannot step on water, stone, or a cell
// that will be flooded at the same second you arrive. Two BFS passes:
// 1) Multi-source BFS from all water to compute flood arrival times.
// 2) BFS from start, only moving to cells reachable before water.
// O(N*M) time, O(N*M) space.

import "fmt"

func minTimeToReachWithoutDrowning(land [][]string) int {
	n, m := len(land), len(land[0])
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	waterTime := make([][]int, n)
	for i := range waterTime {
		waterTime[i] = make([]int, m)
		for j := range waterTime[i] {
			waterTime[i][j] = -1
		}
	}

	var start, dest [2]int
	queue := make([][2]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch land[i][j] {
			case "S":
				start = [2]int{i, j}
			case "D":
				dest = [2]int{i, j}
			case "*":
				waterTime[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// BFS 1: water propagation
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			ni, nj := cur[0]+d[0], cur[1]+d[1]
			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}
			if waterTime[ni][nj] != -1 {
				continue
			}
			if land[ni][nj] == "X" || land[ni][nj] == "D" || land[ni][nj] == "*" {
				continue
			}
			waterTime[ni][nj] = waterTime[cur[0]][cur[1]] + 1
			queue = append(queue, [2]int{ni, nj})
		}
	}

	// BFS 2: player path
	playerTime := make([][]int, n)
	for i := range playerTime {
		playerTime[i] = make([]int, m)
		for j := range playerTime[i] {
			playerTime[i][j] = -1
		}
	}
	playerTime[start[0]][start[1]] = 0
	queue = append(queue, start)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		t := playerTime[cur[0]][cur[1]]

		if cur == dest {
			return t
		}

		for _, d := range dirs {
			ni, nj := cur[0]+d[0], cur[1]+d[1]
			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}
			if playerTime[ni][nj] != -1 {
				continue
			}
			if land[ni][nj] == "X" {
				continue
			}

			nextTime := t + 1
			if land[ni][nj] == "D" {
				playerTime[ni][nj] = nextTime
				queue = append(queue, [2]int{ni, nj})
				continue
			}
			if waterTime[ni][nj] == -1 || nextTime < waterTime[ni][nj] {
				playerTime[ni][nj] = nextTime
				queue = append(queue, [2]int{ni, nj})
			}
		}
	}

	return -1
}

func main() {
	// Example 1: reachable in 3 seconds
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"D", ".", "*"},
		{".", ".", "."},
		{".", "S", "."},
	}))

	// Example 2: blocked by stones -> -1
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"D", "X", "*"},
		{".", ".", "."},
		{".", ".", "S"},
	}))

	// Example 3: longer path
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"D", ".", ".", ".", "*", "."},
		{".", "X", ".", "X", ".", "."},
		{".", ".", ".", ".", "S", "."},
	}))

	// Direct path, no water
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"S", ".", "D"},
	}))

	// No path (stone blocks)
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"S", "X", "D"},
	}))

	// Water cuts off escape
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"S", ".", "D"},
		{"*", ".", "."},
	}))

	// Source is destination
	fmt.Println(minTimeToReachWithoutDrowning([][]string{
		{"D", "."},
		{".", "S"},
	}))
}
