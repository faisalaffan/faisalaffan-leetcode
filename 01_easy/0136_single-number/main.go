package main

// LeetCode #136: Single Number
// https://leetcode.com/problems/single-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func main() {
	fmt.Println(SingleNumber([]int{2, 2, 1}))
	fmt.Println(SingleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(SingleNumber([]int{1}))
}
