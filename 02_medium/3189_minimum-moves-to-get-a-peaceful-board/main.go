package main

// LeetCode #3189: Minimum Moves to Get a Peaceful Board
// https://leetcode.com/problems/minimum-moves-to-get-a-peaceful-board/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minMoves(rooks [][]int) int {
	n := len(rooks)
	rows := make([]int, n)
	cols := make([]int, n)
	for i, r := range rooks {
		rows[i] = r[0]
		cols[i] = r[1]
	}

	sort.Ints(rows)
	sort.Ints(cols)

	moves := 0
	for i := 0; i < n; i++ {
		moves += abs(rows[i] - i)
		moves += abs(cols[i] - i)
	}
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minMoves([][]int{{0, 0}, {1, 1}, {2, 2}}))    // Expected: 0
	fmt.Println(minMoves([][]int{{0, 0}, {0, 2}, {2, 0}}))    // Expected: 2
	fmt.Println(minMoves([][]int{{2, 2}, {0, 0}, {1, 1}}))    // Expected: 0
}
