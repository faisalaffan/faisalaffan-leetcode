package main

// LeetCode #3282: Reach End of Array With Max Score
// https://leetcode.com/problems/reach-end-of-array-with-max-score/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(findMaximumScore([]int{1, 3, 1, 5}))   // 7
	fmt.Println(findMaximumScore([]int{4, 3, 1, 3, 2})) // 16
	fmt.Println(findMaximumScore([]int{2, 2, 2, 2}))    // 6
}

func findMaximumScore(nums []int) int64 {
	var res int64 = 0
	ma := nums[0]
	for i := 1; i < len(nums); i++ {
		res += int64(ma)
		if nums[i] > ma {
			ma = nums[i]
		}
	}
	return res
}
