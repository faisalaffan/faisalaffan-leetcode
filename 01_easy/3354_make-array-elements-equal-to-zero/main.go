package main

// LeetCode #3354: Make Array Elements Equal to Zero
// https://leetcode.com/problems/make-array-elements-equal-to-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MakeArrayElementsEqualToZero([]int{1, 0, 2, 0, 3}))
	fmt.Println(MakeArrayElementsEqualToZero([]int{2, 3, 4, 0, 4, 1, 0}))
}

// MakeArrayElementsEqualToZero counts valid starting positions to make all elements zero.
// From a starting position, move left/right and decrement each visited element by 1.
// Time: O(n^2). Space: O(1).
func MakeArrayElementsEqualToZero(nums []int) int {
	n := len(nums)
	count := 0

	for start := 0; start < n; start++ {
		arr := make([]int, n)
		copy(arr, nums)

		pos := start
		dir := -1 // start going left
		allZero := true
		for _, v := range arr {
			if v != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			count++
			continue
		}

		moves := 0
		for moves < 100000 {
			if arr[pos] > 0 {
				arr[pos]--
			}
			// Check if all zero
			zero := true
			for _, v := range arr {
				if v != 0 {
					zero = false
					break
				}
			}
			if zero {
				count++
				break
			}
			// Move
			nextPos := pos + dir
			if nextPos < 0 || nextPos >= n {
				dir = -dir
			} else {
				pos = nextPos
			}
			moves++
		}
	}
	return count
}
