package main

// LeetCode #1480: Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/
// Difficulty: Easy
//
// LeetCode submission: func runningSum(nums []int) []int

import "fmt"

func main() {
	fmt.Println(RunningSumOfOneDArray([]int{1, 2, 3, 4}))    // [1 3 6 10]
	fmt.Println(RunningSumOfOneDArray([]int{1, 1, 1, 1, 1})) // [1 2 3 4 5]
}

// Time: O(n), Space: O(1) excluding output
func RunningSumOfOneDArray(nums []int) []int {
	res := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		res[i] = sum
	}
	return res
}
