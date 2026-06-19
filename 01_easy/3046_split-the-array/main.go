package main

// LeetCode #3046: Split the Array
// https://leetcode.com/problems/split-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isPossibleToSplit
	fmt.Println(SplitTheArray([]int{1, 1, 2, 2, 3, 4})) // true
	fmt.Println(SplitTheArray([]int{1, 1, 1, 1}))       // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isPossibleToSplit
// Each number can appear at most twice (once in each half of the split)
func SplitTheArray(nums []int) bool {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > 2 {
			return false
		}
	}
	return true
}
