package main

// LeetCode #3000: Maximum Area of Longest Diagonal Rectangle
// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: areaOfMaxDiagonal
	fmt.Println(MaximumAreaOfLongestDiagonalRectangle([][]int{{9, 3}, {8, 6}})) // 48
	fmt.Println(MaximumAreaOfLongestDiagonalRectangle([][]int{{3, 4}, {4, 3}})) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: areaOfMaxDiagonal
func MaximumAreaOfLongestDiagonalRectangle(dimensions [][]int) int {
	maxDiag := 0
	maxArea := 0
	for _, dim := range dimensions {
		l, w := dim[0], dim[1]
		diagSq := l*l + w*w
		area := l * w
		if diagSq > maxDiag || (diagSq == maxDiag && area > maxArea) {
			maxDiag = diagSq
			maxArea = area
		}
	}
	return maxArea
}
