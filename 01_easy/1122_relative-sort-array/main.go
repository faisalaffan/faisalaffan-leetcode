package main

// LeetCode #1122: Relative Sort Array
// https://leetcode.com/problems/relative-sort-array/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	// [2,2,2,1,4,3,3,9,6,7,19]
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
	// [22,28,8,6,17,44]
}

// LeetCode submission: relativeSortArray
func relativeSortArray(arr1, arr2 []int) []int {
	rank := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		rank[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		ri, okI := rank[arr1[i]]
		rj, okJ := rank[arr1[j]]
		if okI && okJ {
			return ri < rj
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
