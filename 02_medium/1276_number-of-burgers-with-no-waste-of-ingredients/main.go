package main

import (
	"fmt"
)

// LeetCode #1276: Number of Burgers with No Waste of Ingredients
// https://leetcode.com/problems/number-of-burgers-with-no-waste-of-ingredients/
// Difficulty: Medium

// Jumbo burger = 4 slices of cheese, 1 tomato. Small = 2 cheese, 1 tomato.
// Given total tomato and cheese slices, find valid (jumbo, small) combination.
// Solve: 4j+2s=cheese, j+s=tomato => j = (cheese - 2*tomato)/2

// Time: O(1)
// Space: O(1)

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// j + s = tomatoSlices
	// 4j + 2s = cheeseSlices
	// => 2j + (2j+2s) = cheeseSlices => 2j + 2*tomatoSlices = cheeseSlices
	// => j = (cheeseSlices - 2*tomatoSlices) / 2

	if cheeseSlices < 0 || tomatoSlices%2 != 0 {
		return []int{}
	}

	jumbo := (cheeseSlices - 2*tomatoSlices) / 2
	small := tomatoSlices - jumbo

	if jumbo < 0 || small < 0 || 4*jumbo+2*small != cheeseSlices {
		return []int{}
	}

	return []int{jumbo, small}
}

func main() {
	fmt.Printf("%v (expected: [1 6])\n", numOfBurgers(16, 7))
	fmt.Printf("%v (expected: [])\n", numOfBurgers(17, 4))
	fmt.Printf("%v (expected: [])\n", numOfBurgers(4, 17))
}
