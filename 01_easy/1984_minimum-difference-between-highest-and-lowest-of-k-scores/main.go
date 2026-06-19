package main

// LeetCode #1984: Minimum Difference Between Highest and Lowest of K Scores
// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{90}, 1))            // 0
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{9, 4, 1, 7}, 2))    // 2
	fmt.Println(MinimumDifferenceBetweenHighestAndLowestOfKScores([]int{9, 4, 1, 7, 5, 3}, 3)) // 3
}

// Time: O(n log n), Space: O(1) ignoring sort
func MinimumDifferenceBetweenHighestAndLowestOfKScores(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := 0; i <= len(nums)-k; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}
