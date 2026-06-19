package main

import (
	"fmt"
)

// LeetCode #1253: Reconstruct a 2-Row Binary Matrix
// https://leetcode.com/problems/reconstruct-a-2-row-binary-matrix/
// Difficulty: Medium

// colsum[i] = sum of col i (top + bottom).
// Fill col with 2 first (top=1, bottom=1), then 1s.
// Greedy: use top row capacity first.

// Time: O(n)
// Space: O(n)

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	top := make([]int, n)
	bottom := make([]int, n)

	for i, s := range colsum {
		if s == 2 {
			top[i] = 1
			bottom[i] = 1
			upper--
			lower--
		}
	}

	if upper < 0 || lower < 0 {
		return [][]int{}
	}

	for i, s := range colsum {
		if s == 1 {
			if upper > 0 {
				top[i] = 1
				upper--
			} else if lower > 0 {
				bottom[i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}

	if upper != 0 || lower != 0 {
		return [][]int{}
	}

	return [][]int{top, bottom}
}

func main() {
	fmt.Printf("%v (expected: [[1 1 0 0] [0 0 1 1]])\n",
		reconstructMatrix(2, 2, []int{1, 1, 1, 1}))

	fmt.Printf("%v (expected: [[]])\n",
		reconstructMatrix(2, 1, []int{1, 1, 1}))

	fmt.Printf("%v\n",
		reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 2}))
}
