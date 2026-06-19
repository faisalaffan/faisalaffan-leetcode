package main

// LeetCode #3424: Minimum Cost to Make Arrays Identical
// https://leetcode.com/problems/minimum-cost-to-make-arrays-identical/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func minCost(arr []int, brr []int, k int64) int64 {
	var cost1 int64
	for i := 0; i < len(arr); i++ {
		diff := arr[i] - brr[i]
		if diff < 0 {
			diff = -diff
		}
		cost1 += int64(diff)
	}

	sortedArr := make([]int, len(arr))
	sortedBrr := make([]int, len(brr))
	copy(sortedArr, arr)
	copy(sortedBrr, brr)
	sort.Ints(sortedArr)
	sort.Ints(sortedBrr)

	var cost2 int64
	for i := 0; i < len(sortedArr); i++ {
		diff := sortedArr[i] - sortedBrr[i]
		if diff < 0 {
			diff = -diff
		}
		cost2 += int64(diff)
	}
	cost2 += k

	if cost1 < cost2 {
		return cost1
	}
	return cost2
}

func main() {
	fmt.Println(minCost([]int{4, 2, 5}, []int{6, 3, 1}, 2)) // 7
	fmt.Println(minCost([]int{1, 2, 3}, []int{4, 5, 6}, 1)) // 9
}
