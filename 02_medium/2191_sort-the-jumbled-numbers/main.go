package main

// LeetCode #2191: Sort the Jumbled Numbers
// https://leetcode.com/problems/sort-the-jumbled-numbers/
// Difficulty: Medium
// Time: O(n log n + n * d) | Space: O(n)

import (
	"fmt"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	type pair struct {
		val    int
		mapped int
		idx    int
	}

	pairs := make([]pair, len(nums))
	for i, num := range nums {
		mapped := 0
		if num == 0 {
			mapped = mapping[0]
		} else {
			digits := []int{}
			n := num
			for n > 0 {
				digits = append(digits, n%10)
				n /= 10
			}
			for j := len(digits) - 1; j >= 0; j-- {
				mapped = mapped*10 + mapping[digits[j]]
			}
		}
		pairs[i] = pair{num, mapped, i}
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].mapped != pairs[j].mapped {
			return pairs[i].mapped < pairs[j].mapped
		}
		return pairs[i].idx < pairs[j].idx
	})

	result := make([]int, len(nums))
	for i, p := range pairs {
		result[i] = p.val
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}))
	// Expected: [338, 38, 991]

	// Test case 2
	fmt.Println(sortJumbled([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}))
	// Expected: [123, 456, 789]
}
