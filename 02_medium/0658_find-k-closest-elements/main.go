package main

// LeetCode #658: Find K Closest Elements
// https://leetcode.com/problems/find-k-closest-elements/
// Difficulty: Medium
// Time: O(log n + k)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
