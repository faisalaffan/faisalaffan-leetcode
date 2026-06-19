package main

// LeetCode #1901: Find a Peak Element II
// https://leetcode.com/problems/find-a-peak-element-ii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindPeakGrid([][]int{{1, 4, 3, 2}, {2, 3, 4, 5}, {5, 4, 3, 6}}))
	fmt.Println(FindPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
}

// Time: O(m log n), Space: O(1)
func FindPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2
		maxRow := 0
		for i := 0; i < m; i++ {
			if mat[i][mid] > mat[maxRow][mid] {
				maxRow = i
			}
		}

		leftVal := -1
		if mid > 0 {
			leftVal = mat[maxRow][mid-1]
		}
		rightVal := -1
		if mid < n-1 {
			rightVal = mat[maxRow][mid+1]
		}

		if mat[maxRow][mid] > leftVal && mat[maxRow][mid] > rightVal {
			return []int{maxRow, mid}
		} else if mat[maxRow][mid] < leftVal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}
