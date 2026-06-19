package main

// LeetCode #287: Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/
// Difficulty: Medium
// Time: O(n), Space: O(1) using Floyd's Cycle Detection

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}))
}
