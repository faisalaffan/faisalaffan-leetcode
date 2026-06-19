package main

// LeetCode #3432: Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{10, 10, 10, 10, 10}))
}

// CountPartitionsWithEvenSumDifference counts partitions where the difference between left and right sums is even.
// Time: O(n). Space: O(1).
func CountPartitionsWithEvenSumDifference(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	leftSum := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum := totalSum - leftSum
		if (leftSum-rightSum)%2 == 0 {
			count++
		}
	}
	return count
}
