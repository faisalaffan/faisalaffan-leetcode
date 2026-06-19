package main

// LeetCode #836: Rectangle Overlap
// https://leetcode.com/problems/rectangle-overlap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3})) // true
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1})) // false
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{2, 2, 3, 3})) // false
}

// isRectangleOverlap checks if two rectangles overlap (positive area).
// Time: O(1). Space: O(1).
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// Check if one rectangle is to the left of the other
	if rec1[2] <= rec2[0] || rec2[2] <= rec1[0] {
		return false
	}
	// Check if one rectangle is above the other
	if rec1[3] <= rec2[1] || rec2[3] <= rec1[1] {
		return false
	}
	return true
}
