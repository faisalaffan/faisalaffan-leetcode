package main

// LeetCode #363: Max Sum of Rectangle No Larger Than K
// https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
// Difficulty: Hard
//
// For each pair of rows (top, bottom), compute column sums and apply a
// 1D Kadane-with-bound approach: maintain a sorted prefix-sum list and
// binary-search for the smallest prefix >= current - k.
// O(m^2 * n * log n) time, O(n) space.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1: [[1,0,1],[0,-2,3]], k=2 -> 2 ([[0,-2],[0,3]] sum=2)
	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	// Example 2: [[2,2,-1]], k=3 -> 3
	fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
	// Example 3: [[1]], k=1 -> 1
	fmt.Println(maxSumSubmatrix([][]int{{1}}, 0))
	// Edge: larger matrix
	fmt.Println(maxSumSubmatrix([][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 10))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	best := math.MinInt64

	for top := 0; top < rows; top++ {
		colSums := make([]int, cols)
		for bottom := top; bottom < rows; bottom++ {
			// Accumulate row bottom into column sums
			for c := 0; c < cols; c++ {
				colSums[c] += matrix[bottom][c]
			}

			// 1D max subarray sum no larger than k on colSums
			best = max(best, maxSumNoLargerThanK(colSums, k))
			if best == k {
				return k
			}
		}
	}
	return best
}

// maxSumNoLargerThanK finds max subarray sum <= k using sorted prefix sums.
func maxSumNoLargerThanK(arr []int, k int) int {
	prefix := 0
	best := math.MinInt64

	// Sorted prefix sums (we maintain a sorted list)
	sorted := []int{0}

	for _, v := range arr {
		prefix += v
		// Find smallest prefix in sorted >= prefix - k
		idx := sort.SearchInts(sorted, prefix-k)
		if idx < len(sorted) {
			best = max(best, prefix-sorted[idx])
		}
		// Insert prefix into sorted position
		ins := sort.SearchInts(sorted, prefix)
		sorted = append(sorted, 0)
		copy(sorted[ins+1:], sorted[ins:])
		sorted[ins] = prefix
	}

	return best
}
