package main

// LeetCode #2966: Divide Array Into Arrays With Max Difference
// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divideArray2966([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray2966([]int{1, 3, 3, 2, 7, 3}, 3))
	fmt.Println(divideArray2966([]int{1, 2, 3}, 0))
}

func divideArray2966(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i += 3 {
		t := make([]int, 3)
		copy(t, nums[i:i+3])
		if t[2]-t[0] > k {
			return [][]int{}
		}
		ans = append(ans, t)
	}
	return ans
}
