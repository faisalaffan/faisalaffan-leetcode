package main

// LeetCode #1007: Minimum Domino Rotations For Equal Row
// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
// Difficulty: Medium
//
// Approach: Check if we can make all values equal to tops[0] or bottoms[0]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2})) // 2
	fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))       // -1
	fmt.Println(minDominoRotations([]int{1, 2, 1, 1, 1, 2, 2, 2}, []int{2, 1, 2, 2, 2, 1, 1, 1})) // 2
}

func minDominoRotations(tops []int, bottoms []int) int {
	result := check(tops, bottoms, tops[0])
	if result != -1 {
		return result
	}
	return check(tops, bottoms, bottoms[0])
}

func check(tops, bottoms []int, target int) int {
	topRot := 0
	bottomRot := 0

	for i := 0; i < len(tops); i++ {
		if tops[i] != target && bottoms[i] != target {
			return -1
		}
		if tops[i] != target {
			topRot++
		}
		if bottoms[i] != target {
			bottomRot++
		}
	}

	if topRot < bottomRot {
		return topRot
	}
	return bottomRot
}
