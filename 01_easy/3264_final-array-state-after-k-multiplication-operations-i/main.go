package main

// LeetCode #3264: Final Array State After K Multiplication Operations I
// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FinalArrayStateAfterKMultiplicationOperationsI([]int{2, 1, 3, 5, 6}, 5, 2))
	fmt.Println(FinalArrayStateAfterKMultiplicationOperationsI([]int{1, 2}, 3, 4))
}

// FinalArrayStateAfterKMultiplicationOperationsI finds the minimum element each time, multiplies it by multiplier, and repeats k times.
// Time: O(k * n). Space: O(1).
func FinalArrayStateAfterKMultiplicationOperationsI(nums []int, k int, multiplier int) []int {
	for t := 0; t < k; t++ {
		// Find index of minimum element
		minIdx := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[minIdx] *= multiplier
	}
	return nums
}
