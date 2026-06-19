package main

// LeetCode #2626: Array Reduce Transformation
// https://leetcode.com/problems/array-reduce-transformation/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Reduces slice using a function.

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := func(acc, curr int) int { return acc + curr }
	fmt.Println(arrayReduceTransformation(nums, sum, 0))

	nums2 := []int{1, 2, 3, 4}
	product := func(acc, curr int) int { return acc * curr }
	fmt.Println(arrayReduceTransformation(nums2, product, 1))
}

func arrayReduceTransformation(nums []int, fn func(int, int) int, init int) int {
	result := init
	for _, num := range nums {
		result = fn(result, num)
	}
	return result
}
