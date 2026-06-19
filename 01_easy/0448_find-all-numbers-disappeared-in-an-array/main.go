package main

// LeetCode #448: Find All Numbers Disappeared in an Array
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindAllNumbersDisappearedInAnArray(nums []int) []int {
	for _, v := range nums {
		idx := v
		if idx < 0 {
			idx = -idx
		}
		idx--
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	var result []int
	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	fmt.Println(FindAllNumbersDisappearedInAnArray([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(FindAllNumbersDisappearedInAnArray([]int{1, 1}))
}
