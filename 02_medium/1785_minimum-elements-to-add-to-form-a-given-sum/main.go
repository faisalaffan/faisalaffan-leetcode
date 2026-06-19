package main

// LeetCode #1785: Minimum Elements to Add to Form a Given Sum
// https://leetcode.com/problems/minimum-elements-to-add-to-form-a-given-sum/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	diff := goal - sum
	if diff < 0 {
		diff = -diff
	}

	// Minimum elements = ceil(diff / limit)
	return (diff + limit - 1) / limit
}

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4)) // Expected: 2
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0)) // Expected: 1
	fmt.Println(minElements([]int{0}, 1, 1000000)) // Expected: 1000000
}
