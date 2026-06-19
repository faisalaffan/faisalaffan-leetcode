package main

// LeetCode #977: Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4,9,9,49,121]
}

// sortedSquares returns squares of each number sorted in non-decreasing order.
// Time: O(n). Space: O(n).
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	l, r := 0, n-1
	pos := n - 1
	for l <= r {
		leftSq := nums[l] * nums[l]
		rightSq := nums[r] * nums[r]
		if leftSq > rightSq {
			result[pos] = leftSq
			l++
		} else {
			result[pos] = rightSq
			r--
		}
		pos--
	}
	return result
}
