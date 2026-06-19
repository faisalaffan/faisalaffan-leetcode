package main

// LeetCode #2535: Difference Between Element Sum and Digit Sum of an Array
// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/
// Difficulty: Easy
// Time O(n log m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DifferenceBetweenElementSumAndDigitSumOfAnArray([]int{1, 15, 6, 3})) // 9
	fmt.Println(DifferenceBetweenElementSumAndDigitSumOfAnArray([]int{1, 2, 3, 4}))  // 0
}

func DifferenceBetweenElementSumAndDigitSumOfAnArray(nums []int) int {
	elementSum := 0
	digitSum := 0
	for _, n := range nums {
		elementSum += n
		for n > 0 {
			digitSum += n % 10
			n /= 10
		}
	}
	result := elementSum - digitSum
	if result < 0 {
		return -result
	}
	return result
}
