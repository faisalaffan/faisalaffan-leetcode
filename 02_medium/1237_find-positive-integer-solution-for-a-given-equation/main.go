package main

import (
	"fmt"
)

// LeetCode #1237: Find Positive Integer Solution for a Given Equation
// https://leetcode.com/problems/find-positive-integer-solution-for-a-given-equation/
// Difficulty: Medium

// Given custom function f(x,y) that's monotonically increasing in both x and y,
// find all [x,y] where f(x,y) = z.
// Start from x=1, y=1000 and move inwards.

// Time: O(x + y) where x,y in [1, 1000]
// Space: O(n) for result

type customFunction func(int, int) int

func findSolution(customfunction customFunction, z int) [][]int {
	result := make([][]int, 0)
	x, y := 1, 1000

	for x <= 1000 && y >= 1 {
		val := customfunction(x, y)
		if val == z {
			result = append(result, []int{x, y})
			x++
			y--
		} else if val < z {
			x++
		} else {
			y--
		}
	}

	return result
}

func main() {
	// f(x,y) = x + y
	f1 := func(x, y int) int { return x + y }
	fmt.Printf("%v (expected: [[1 4] [2 3] [3 2] [4 1]])\n", findSolution(f1, 5))

	// f(x,y) = x * y
	f2 := func(x, y int) int { return x * y }
	fmt.Printf("%v (expected: [[1 12] [2 6] [3 4] [4 3] [6 2] [12 1]])\n", findSolution(f2, 12))
}
