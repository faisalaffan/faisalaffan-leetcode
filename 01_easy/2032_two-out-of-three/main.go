package main

// LeetCode #2032: Two Out of Three
// https://leetcode.com/problems/two-out-of-three/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TwoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))    // [3 2]
	fmt.Println(TwoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2}))       // [2 3 1]
	fmt.Println(TwoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5}))    // []
}

// Time: O(n), Space: O(n)
func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	set3 := make(map[int]bool)

	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	for _, v := range nums3 {
		set3[v] = true
	}

	count := make(map[int]int)
	for v := range set1 {
		count[v]++
	}
	for v := range set2 {
		count[v]++
	}
	for v := range set3 {
		count[v]++
	}

	var result []int
	for v, c := range count {
		if c >= 2 {
			result = append(result, v)
		}
	}
	sort.Ints(result)
	return result
}
