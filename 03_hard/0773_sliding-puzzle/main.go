package main

// LeetCode #773: Sliding Puzzle
// https://leetcode.com/problems/sliding-puzzle/
// Difficulty: Hard
//
// 2x3 board with tiles 1-5 and 0 (empty). Find minimum number of
// moves to reach [[1,2,3],[4,5,0]]. If impossible, return -1.
//
// Approach: BFS
// Convert board to string "123450", BFS from start state, tracking
// visited states. Allowed moves: 0 can swap with adjacent positions.

import "fmt"

func main() {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}})) // 1
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}})) // -1
	fmt.Println(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}})) // 5
	fmt.Println(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}})) // 14
}

func slidingPuzzle(board [][]int) int {
	// Convert board to string
	start := ""
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			start += string(rune('0' + board[i][j]))
		}
	}

	target := "123450"

	// Neighbors for each position in the 1D string (index 0-5)
	// Positions on the 2x3 board:
	// 0 1 2
	// 3 4 5
	neighbors := [][]int{
		{1, 3},       // 0: right, down
		{0, 2, 4},    // 1: left, right, down
		{1, 5},       // 2: left, down
		{0, 4},       // 3: up, right
		{1, 3, 5},    // 4: up, left, right
		{2, 4},       // 5: up, left
	}

	if start == target {
		return 0
	}

	visited := map[string]bool{start: true}
	queue := []string{start}
	depth := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]

			// Find position of '0'
			zeroIdx := 0
			for cur[zeroIdx] != '0' {
				zeroIdx++
			}

			// Try all neighbors of '0'
			for _, nb := range neighbors[zeroIdx] {
				next := []byte(cur)
				next[zeroIdx], next[nb] = next[nb], next[zeroIdx]
				nextStr := string(next)

				if nextStr == target {
					return depth + 1
				}
				if !visited[nextStr] {
					visited[nextStr] = true
					queue = append(queue, nextStr)
				}
			}
		}
		queue = queue[size:]
		depth++
	}

	return -1
}
