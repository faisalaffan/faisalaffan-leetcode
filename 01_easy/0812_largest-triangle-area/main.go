package main

// LeetCode #812: Largest Triangle Area
// https://leetcode.com/problems/largest-triangle-area/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}})) // 2.0
	fmt.Println(largestTriangleArea([][]int{{1, 0}, {0, 0}, {0, 1}}))                 // 0.5
}

// largestTriangleArea finds the largest area of any triangle formed by any 3 points.
// Time: O(n^3). Space: O(1).
func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := areaTriangle(points[i], points[j], points[k])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func areaTriangle(a, b, c []int) float64 {
	return math.Abs(float64(a[0]*(b[1]-c[1])+b[0]*(c[1]-a[1])+c[0]*(a[1]-b[1]))) / 2.0
}
