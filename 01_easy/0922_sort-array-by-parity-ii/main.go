package main

// LeetCode #922: Sort Array By Parity II
// https://leetcode.com/problems/sort-array-by-parity-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7})) // [4,5,2,7] or [4,7,2,5]
	fmt.Println(sortArrayByParityII([]int{2, 3}))        // [2,3]
}

// sortArrayByParityII puts even numbers at even indices, odd numbers at odd indices.
// Time: O(n). Space: O(1).
func sortArrayByParityII(nums []int) []int {
	j := 1 // odd pointer
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
