package main

// LeetCode #704: Binary Search
// https://leetcode.com/problems/binary-search/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))    // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))    // -1
	fmt.Println(search([]int{5}, 5))                      // 0
}

// search performs binary search on a sorted array.
// Time: O(log n). Space: O(1).
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
