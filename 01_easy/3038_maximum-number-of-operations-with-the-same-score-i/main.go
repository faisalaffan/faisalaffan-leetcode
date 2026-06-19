package main

// LeetCode #3038: Maximum Number of Operations With the Same Score I
// https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maxOperations
	fmt.Println(MaximumNumberOfOperationsWithTheSameScoreI([]int{3, 2, 1, 4, 5})) // 2
	fmt.Println(MaximumNumberOfOperationsWithTheSameScoreI([]int{3, 2, 6, 1, 4})) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maxOperations
func MaximumNumberOfOperationsWithTheSameScoreI(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	score := nums[0] + nums[1]
	count := 1
	for i := 2; i+1 < len(nums); i += 2 {
		if nums[i]+nums[i+1] == score {
			count++
		} else {
			break
		}
	}
	return count
}
