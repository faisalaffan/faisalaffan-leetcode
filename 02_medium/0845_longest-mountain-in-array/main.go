package main

// LeetCode #845: Longest Mountain in Array
// https://leetcode.com/problems/longest-mountain-in-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LongestMountainInArray([]int{2, 1, 4, 7, 3, 2, 5}))
	fmt.Println(LongestMountainInArray([]int{2, 2, 2}))
	fmt.Println(LongestMountainInArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

// Time: O(n) | Space: O(1)
func LongestMountainInArray(arr []int) int {
	n := len(arr)
	ans := 0
	i := 1

	for i < n {
		// Skip non-increasing start
		for i < n && arr[i] == arr[i-1] {
			i++
		}

		// Climb up
		up := 0
		for i < n && arr[i] > arr[i-1] {
			up++
			i++
		}

		// Climb down
		down := 0
		for i < n && arr[i] < arr[i-1] {
			down++
			i++
		}

		// Valid mountain needs both up and down segments
		if up > 0 && down > 0 {
			if up+down+1 > ans {
				ans = up + down + 1
			}
		}
	}

	return ans
}
