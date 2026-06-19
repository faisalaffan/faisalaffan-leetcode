package main

// LeetCode #2574: Left and Right Sum Differences
// https://leetcode.com/problems/left-and-right-sum-differences/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(LeftAndRightSumDifferences([]int{10, 4, 8, 3})) // [15,1,11,22]
	fmt.Println(LeftAndRightSumDifferences([]int{1}))            // [0]
}

func LeftAndRightSumDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	res := make([]int, n)
	leftSum := 0
	for i, v := range nums {
		rightSum := total - leftSum - v
		diff := leftSum - rightSum
		if diff < 0 {
			diff = -diff
		}
		res[i] = diff
		leftSum += v
	}
	return res
}
