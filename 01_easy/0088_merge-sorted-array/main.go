package main

// LeetCode #88: Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/
// Difficulty: Easy

import "fmt"

// Time: O(m+n) | Space: O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	Merge(n1, 3, []int{2, 5, 6}, 3)
	fmt.Println(n1)
	n2 := []int{1}
	Merge(n2, 1, []int{}, 0)
	fmt.Println(n2)
}
