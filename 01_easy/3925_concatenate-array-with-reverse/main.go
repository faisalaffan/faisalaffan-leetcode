package main

// LeetCode #3925: Concatenate Array With Reverse
// https://leetcode.com/problems/concatenate-array-with-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenateArrayWithReverse([]int{1, 2, 3}))
	fmt.Println(ConcatenateArrayWithReverse([]int{1}))
}

// Time: O(n)
// Space: O(n)
func ConcatenateArrayWithReverse(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = nums[n-1-i]
	}
	return ans
}
