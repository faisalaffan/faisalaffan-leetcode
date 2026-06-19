package main

// LeetCode #3683: Earliest Time to Finish One Task
// https://leetcode.com/problems/earliest-time-to-finish-one-task/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(EarliestTimeToFinishOneTask([][]int{{1, 6}, {2, 3}}))
	fmt.Println(EarliestTimeToFinishOneTask([][]int{{100, 100}, {100, 100}, {100, 100}}))
}

// Time: O(n)
// Space: O(1)
func EarliestTimeToFinishOneTask(tasks [][]int) int {
	ans := math.MaxInt
	for _, t := range tasks {
		finish := t[0] + t[1]
		if finish < ans {
			ans = finish
		}
	}
	return ans
}
