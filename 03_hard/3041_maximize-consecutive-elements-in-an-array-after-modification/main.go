package main

// LeetCode #3041: Maximize Consecutive Elements in an Array After Modification
// https://leetcode.com/problems/maximize-consecutive-elements-in-an-array-after-modification/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)
	f := make(map[int]int)
	ans := 0
	for _, x := range nums {
		f[x] = max2(f[x], f[x-1]+1)
		f[x+1] = max2(f[x+1], f[x]+1)
	}
	for _, v := range f {
		if v > ans { ans = v }
	}
	return ans
}
func max2(a, b int) int { if a > b { return a }; return b }

func main() {
	fmt.Println(maxSelectedElements([]int{2, 1, 5, 1, 1}))
	fmt.Println(maxSelectedElements([]int{1, 4, 7, 10}))
}
