package main

// LeetCode #1464: Maximum Product of Two Elements in an Array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func maxProduct(nums []int) int

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 4, 5, 2})) // 12
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{1, 5, 4, 5})) // 16
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 7}))       // 12
}

// Time: O(n), Space: O(1)
func MaximumProductOfTwoElementsInAnArray(nums []int) int {
	first, second := 0, 0
	for _, v := range nums {
		if v > first {
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}
	return (first - 1) * (second - 1)
}
