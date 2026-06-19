package main

// LeetCode #1539: Kth Missing Positive Number
// https://leetcode.com/problems/kth-missing-positive-number/
// Difficulty: Easy
//
// LeetCode submission: func findKthPositive(arr []int, k int) int

import "fmt"

func main() {
	fmt.Println(KthMissingPositiveNumber([]int{2, 3, 4, 7, 11}, 5)) // 9
	fmt.Println(KthMissingPositiveNumber([]int{1, 2, 3, 4}, 2))     // 6
	fmt.Println(KthMissingPositiveNumber([]int{1, 3, 5}, 2))        // 4
}

// Time: O(log n), Space: O(1)
func KthMissingPositiveNumber(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid]-mid-1 < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo + k
}
