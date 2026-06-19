package main

// LeetCode #163: Missing Ranges
// https://leetcode.com/problems/missing-ranges/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"strconv"
)

// Time: O(n) | Space: O(1) excluding output
func FindMissingRanges(nums []int, lower int, upper int) []string {
	var res []string
	addRange := func(lo, hi int) {
		if lo > hi {
			return
		}
		if lo == hi {
			res = append(res, strconv.Itoa(lo))
		} else {
			res = append(res, strconv.Itoa(lo)+"->"+strconv.Itoa(hi))
		}
	}
	prev := lower - 1
	for i := 0; i <= len(nums); i++ {
		var curr int
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}
		if curr-prev > 1 {
			addRange(prev+1, curr-1)
		}
		prev = curr
	}
	return res
}

func main() {
	fmt.Println(FindMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
	fmt.Println(FindMissingRanges([]int{-1}, -1, -1))
}
