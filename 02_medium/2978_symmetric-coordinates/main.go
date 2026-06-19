package main

// LeetCode #2978: Symmetric Coordinates
// https://leetcode.com/problems/symmetric-coordinates/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Coordinate struct {
	X int
	Y int
}

func symmetricCoordinates(coords []Coordinate) []Coordinate {
	// Build set of all coordinates
	coordSet := make(map[[2]int]bool)
	for _, c := range coords {
		coordSet[[2]int{c.X, c.Y}] = true
	}

	seen := make(map[[2]int]bool)
	var results []Coordinate

	for _, c := range coords {
		x, y := c.X, c.Y
		if x > y {
			continue
		}

		if seen[[2]int{x, y}] {
			continue
		}
		seen[[2]int{x, y}] = true

		// Check if symmetric pair (y,x) also exists
		if coordSet[[2]int{y, x}] {
			results = append(results, Coordinate{X: x, Y: y})
		}
	}

	// Order by X ASC, Y ASC
	sort.Slice(results, func(i, j int) bool {
		if results[i].X != results[j].X {
			return results[i].X < results[j].X
		}
		return results[i].Y < results[j].Y
	})

	return results
}

func main() {
	coords := []Coordinate{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
		{X: 2, Y: 1},
		{X: 5, Y: 6},
		{X: 4, Y: 3},
		{X: 7, Y: 7},
		{X: 8, Y: 9},
		{X: 1, Y: 3},
	}

	fmt.Println("Symmetric Coordinates")
	fmt.Println("====================")
	fmt.Printf("%-6s %s\n", "X", "Y")
	fmt.Println("-----------")

	results := symmetricCoordinates(coords)
	for _, r := range results {
		fmt.Printf("%-6d %d\n", r.X, r.Y)
	}
}
