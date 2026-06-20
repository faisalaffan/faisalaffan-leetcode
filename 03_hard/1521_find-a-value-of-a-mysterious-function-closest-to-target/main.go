package main

// LeetCode #1521: Find a Value of a Mysterious Function Closest to Target
// https://leetcode.com/problems/find-a-value-of-a-mysterious-function-closest-to-target/
// Difficulty: Hard
//
// Winston has a mysterious function func(arr, l, r) that returns the
// bitwise AND of all elements in arr[l..r]. Find the minimum absolute
// difference between any func value and target.
//
// Approach: Track all possible AND values of subarrays ending at each
// position. AND values only decrease, so the set of distinct values
// is small (at most 32 per position).

import "fmt"

func main() {
	// Example 1
	fmt.Println(closestToTarget([]int{9, 12, 3, 7, 15}, 5))
	// Example 2
	fmt.Println(closestToTarget([]int{1000000, 1000000, 1000000}, 1))
	// Edge: single element
	fmt.Println(closestToTarget([]int{5}, 5))
}

func closestToTarget(arr []int, target int) int {
	ans := abs(arr[0] - target)
	pre := map[int]bool{arr[0]: true}
	for _, x := range arr {
		cur := map[int]bool{x: true}
		for y := range pre {
			cur[x&y] = true
		}
		for y := range cur {
			ans = min(ans, abs(y-target))
		}
		pre = cur
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
