package main

// LeetCode #2824: Count Pairs Whose Sum is Less than Target
// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}

func CountPairsWhoseSumIsLessThanTarget(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}
