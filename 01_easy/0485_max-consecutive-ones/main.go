package main

// LeetCode #485: Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}

func main() {
	fmt.Println(MaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(MaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
