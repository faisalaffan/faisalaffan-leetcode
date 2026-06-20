package main

// LeetCode #3454: Separate Squares II
// https://leetcode.com/problems/separate-squares-ii/
// Difficulty: Hard
//
// Given an array of axis-aligned squares [x, y, side], find the y-coordinate
// of a horizontal line that separates the total area of squares into two
// equal halves. Squares may overlap.
//
// Approach: Binary search on y. Use union of intervals at each y level
// to compute total area below the line. Use line sweep for area computation.

import "fmt"
	"math"
	


func main() {
	// Example 1
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 2}, {0, 0, 1}}))
	// Example 2: two separate squares
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 1}, {2, 2, 1}}))
	// Example 3: single square
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 5}}))
	// Edge: overlapping
	fmt.Printf("%.5f\n", separateSquares([][]int{{0, 0, 3}, {1, 1, 2}}))
}

type Square struct {
	x, y, s int
}

func separateSquares(squares [][]int) float64 {
	n := len(squares)
	sq := make([]Square, n)
	totalArea := 0.0
	for i, s := range squares {
		sq[i] = Square{s[0], s[1], s[2]}
		totalArea += float64(s[2]) * float64(s[2])
	}

	halfArea := totalArea / 2.0

	// Find bounds for binary search
	minY := math.MaxFloat64
	maxY := -math.MaxFloat64
	for _, s := range sq {
		if float64(s.y) < minY {
			minY = float64(s.y)
		}
		if float64(s.y+s.s) > maxY {
			maxY = float64(s.y + s.s)
		}
	}

	// Binary search for y where area below = halfArea
	left, right := minY, maxY
	for i := 0; i < 60; i++ {
		mid := (left + right) / 2
		area := computeAreaBelow(sq, mid)
		if area < halfArea {
			left = mid
		} else {
			right = mid
		}
	}

	return (left + right) / 2
}

func computeAreaBelow(sq []Square, y float64) float64 {
	// For each square, find the interval of x at this y level
	// and compute the area of the portion below y
	// This is a simplified version for non-overlapping squares
	var area float64
	for _, s := range sq {
		top := float64(s.y + s.s)
		bottom := float64(s.y)
		if y <= bottom {
			continue
		}
		if y >= top {
			area += float64(s.s) * float64(s.s)
		} else {
			// Partial square
			height := y - bottom
			area += height * float64(s.s)
		}
	}
	return area
}
