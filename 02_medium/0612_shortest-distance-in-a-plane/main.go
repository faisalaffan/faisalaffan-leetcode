package main

// LeetCode #612: Shortest Distance in a Plane
// https://leetcode.com/problems/shortest-distance-in-a-plane/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	points := [][]float64{
		{-1, -1},
		{0, 0},
		{1, 1},
		{2, 2},
	}
	fmt.Printf("%.4f\n", ShortestDistance(points))
}

func ShortestDistance(points [][]float64) float64 {
	if len(points) < 2 {
		return 0
	}

	minDist := math.MaxFloat64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}
