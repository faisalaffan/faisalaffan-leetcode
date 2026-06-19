package main

// LeetCode #1232: Check If It Is a Straight Line
// https://leetcode.com/problems/check-if-it-is-a-straight-line/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}})) // true
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}})) // false
}

// LeetCode submission: checkStraightLine
func checkStraightLine(coordinates [][]int) bool {
	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]
	dx, dy := x1-x0, y1-y0
	for i := 2; i < len(coordinates); i++ {
		xi, yi := coordinates[i][0], coordinates[i][1]
		if (xi-x0)*dy != (yi-y0)*dx {
			return false
		}
	}
	return true
}
