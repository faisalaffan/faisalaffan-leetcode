package main

// LeetCode #2850: Minimum Moves to Spread Stones Over Grid
// https://leetcode.com/problems/minimum-moves-to-spread-stones-over-grid/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumMovesToSpreadStonesOverGrid(grid [][]int) int {
	// Grid is always 3x3
	// Collect positions of empty cells (0) and cells with extra stones (>1)
	zeros := make([][2]int, 0)
	extras := make([][2]int, 0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				zeros = append(zeros, [2]int{i, j})
			}
			for k := 1; k < grid[i][j]; k++ {
				extras = append(extras, [2]int{i, j})
			}
		}
	}

	if len(zeros) == 0 {
		return 0
	}

	// Use permutation to find minimum total moves
	n := len(zeros)
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}

	best := math.MaxInt32
	for {
		total := 0
		for i, p := range perm {
			dx := extras[i][0] - zeros[p][0]
			if dx < 0 {
				dx = -dx
			}
			dy := extras[i][1] - zeros[p][1]
			if dy < 0 {
				dy = -dy
			}
			total += dx + dy
		}
		if total < best {
			best = total
		}

		// Next permutation
		k := -1
		for i := n - 2; i >= 0; i-- {
			if perm[i] < perm[i+1] {
				k = i
				break
			}
		}
		if k == -1 {
			break
		}
		l := n - 1
		for perm[l] <= perm[k] {
			l--
		}
		perm[k], perm[l] = perm[l], perm[k]
		for i, j := k+1, n-1; i < j; i, j = i+1, j-1 {
			perm[i], perm[j] = perm[j], perm[i]
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumMovesToSpreadStonesOverGrid([][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}}))
	fmt.Println(MinimumMovesToSpreadStonesOverGrid([][]int{{1, 3, 0}, {1, 0, 0}, {1, 0, 3}}))
}
