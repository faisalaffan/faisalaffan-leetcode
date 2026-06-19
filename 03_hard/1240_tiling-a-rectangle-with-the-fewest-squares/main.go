package main

// LeetCode #1240: Tiling a Rectangle with the Fewest Squares
// https://leetcode.com/problems/tiling-a-rectangle-with-the-fewest-squares/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("1240. Tiling a Rectangle with the Fewest Squares")
	fmt.Println("n=2, m=3:", tilingRectangle(2, 3), "(expected 3)")
	fmt.Println("n=5, m=8:", tilingRectangle(5, 8), "(expected 5)")
	fmt.Println("n=3, m=3:", tilingRectangle(3, 3), "(expected 1)")
}

func tilingRectangle(n int, m int) int {
	// Ensure n >= m for column-minimization
	if n < m {
		n, m = m, n
	}

	ans := n * m // worst case: all 1x1
	height := make([]int, m)

	var dfs func(used int)
	dfs = func(used int) {
		if used >= ans {
			return
		}

		// Find the lowest (minimum height) column; pick the leftmost one.
		minRow := math.MaxInt32
		col := -1
		for j := 0; j < m; j++ {
			if height[j] < minRow {
				minRow = height[j]
				col = j
			}
		}

		// All columns filled to the top.
		if minRow == n {
			if used < ans {
				ans = used
			}
			return
		}

		// Compute max possible square size at (col, minRow).
		maxSize := 1
		for col+maxSize <= m && minRow+maxSize <= n {
			// Check if all positions from col to col+maxSize-1 have height == minRow.
			valid := true
			for j := col; j < col+maxSize; j++ {
				if height[j] != minRow {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
			maxSize++
		}
		maxSize--

		// Try sizes from largest to smallest (better pruning).
		for size := maxSize; size >= 1; size-- {
			for j := col; j < col+size; j++ {
				height[j] += size
			}
			dfs(used + 1)
			for j := col; j < col+size; j++ {
				height[j] -= size
			}
		}
	}

	dfs(0)
	return ans
}
