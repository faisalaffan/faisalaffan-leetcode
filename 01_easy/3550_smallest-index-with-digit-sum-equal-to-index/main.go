package main

// LeetCode #3550: Smallest Index With Digit Sum Equal to Index
// https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestIndexWithDigitSumEqualToIndex([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(SmallestIndexWithDigitSumEqualToIndex([]int{10, 11, 12, 13, 14}))
}

// digitSum returns sum of digits.
func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// SmallestIndexWithDigitSumEqualToIndex returns the smallest index i where digit sum of nums[i] equals i, else -1.
// Time: O(n * log max). Space: O(1).
func SmallestIndexWithDigitSumEqualToIndex(nums []int) int {
	for i, v := range nums {
		if sumDigits(v) == i {
			return i
		}
	}
	return -1
}
