package main

// LeetCode #1640: Check Array Formation Through Concatenation
// https://leetcode.com/problems/check-array-formation-through-concatenation/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CanFormArray(arr []int, pieces [][]int) bool {
	pos := make(map[int]int)
	for i, num := range arr {
		pos[num] = i
	}
	for _, piece := range pieces {
		first := piece[0]
		idx, ok := pos[first]
		if !ok {
			return false
		}
		for j := 1; j < len(piece); j++ {
			if idx+j >= len(arr) || arr[idx+j] != piece[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(CanFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(CanFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	fmt.Println(CanFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}
