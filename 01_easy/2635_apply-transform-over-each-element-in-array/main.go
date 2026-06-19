package main

// LeetCode #2635: Apply Transform Over Each Element in Array
// https://leetcode.com/problems/apply-transform-over-each-element-in-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Maps a function over a slice.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	double := func(n int, i int) int { return n * 2 }
	fmt.Println(ApplyTransformOverEachElementInArray(nums, double))

	nums2 := []int{1, 2, 3}
	timesIndex := func(n int, i int) int { return n * i }
	fmt.Println(ApplyTransformOverEachElementInArray(nums2, timesIndex))
}

func ApplyTransformOverEachElementInArray(arr []int, fn func(int, int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v, i)
	}
	return result
}
