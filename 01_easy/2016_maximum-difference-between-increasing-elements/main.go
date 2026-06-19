package main

// LeetCode #2016: Maximum Difference Between Increasing Elements
// https://leetcode.com/problems/maximum-difference-between-increasing-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{7, 1, 5, 4}))   // 4
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{9, 4, 3, 2}))   // -1
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{1, 5, 2, 10}))  // 9
}

// Time: O(n), Space: O(1)
func MaximumDifferenceBetweenIncreasingElements(nums []int) int {
	minSoFar := nums[0]
	maxDiff := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > minSoFar {
			diff := nums[i] - minSoFar
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		if nums[i] < minSoFar {
			minSoFar = nums[i]
		}
	}
	return maxDiff
}
