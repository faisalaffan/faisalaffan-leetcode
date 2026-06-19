package main

// LeetCode #3061: Calculate Trapping Rain Water (SQL simulation)
// https://leetcode.com/problems/calculate-trapping-rain-water/
// Difficulty: Hard [Paid]

import "fmt"

type Height struct {
	ID     int
	Height int
}

func calculateTrappingRainWater(heights []Height) int {
	n := len(heights)
	if n == 0 { return 0 }
	h := make([]int, n)
	for i, ht := range heights { h[i] = ht.Height }
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = h[0]
	for i := 1; i < n; i++ {
		if h[i] > leftMax[i-1] { leftMax[i] = h[i] } else { leftMax[i] = leftMax[i-1] }
	}
	rightMax[n-1] = h[n-1]
	for i := n - 2; i >= 0; i-- {
		if h[i] > rightMax[i+1] { rightMax[i] = h[i] } else { rightMax[i] = rightMax[i+1] }
	}
	total := 0
	for i := 0; i < n; i++ {
		minBound := leftMax[i]
		if rightMax[i] < minBound { minBound = rightMax[i] }
		total += minBound - h[i]
	}
	return total
}

func main() {
	heights := []Height{
		{1, 0}, {2, 1}, {3, 0}, {4, 2},
		{5, 1}, {6, 0}, {7, 1}, {8, 3},
		{9, 2}, {10, 1}, {11, 2}, {12, 1},
	}
	fmt.Println(calculateTrappingRainWater(heights))
}
