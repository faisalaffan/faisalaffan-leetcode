package main

// LeetCode #912: Sort an Array
// https://leetcode.com/problems/sort-an-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SortAnArray([]int{5, 2, 3, 1}))
	fmt.Println(SortAnArray([]int{5, 1, 1, 2, 0, 0}))
	fmt.Println(SortAnArray([]int{3, -1}))
}

// Time: O(n log n) | Space: O(n)
func SortAnArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := SortAnArray(nums[:mid])
	right := SortAnArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
