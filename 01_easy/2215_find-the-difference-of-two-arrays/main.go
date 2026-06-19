package main

// LeetCode #2215: Find the Difference of Two Arrays
// https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheDifferenceOfTwoArrays([]int{1, 2, 3}, []int{2, 4, 6}))       // [[1 3] [4 6]]
	fmt.Println(FindTheDifferenceOfTwoArrays([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})) // [[3] []]
}

// Time: O(n + m), Space: O(n + m)
func FindTheDifferenceOfTwoArrays(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}

	var diff1 []int
	for v := range set1 {
		if !set2[v] {
			diff1 = append(diff1, v)
		}
	}

	var diff2 []int
	for v := range set2 {
		if !set1[v] {
			diff2 = append(diff2, v)
		}
	}

	return [][]int{diff1, diff2}
}
