package main

// LeetCode #2206: Divide Array Into Equal Pairs
// https://leetcode.com/problems/divide-array-into-equal-pairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideArrayIntoEqualPairs([]int{3, 2, 3, 2, 2, 2})) // true
	fmt.Println(DivideArrayIntoEqualPairs([]int{1, 2, 3, 4}))       // false
}

// Time: O(n), Space: O(n)
func DivideArrayIntoEqualPairs(nums []int) bool {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, c := range freq {
		if c%2 != 0 {
			return false
		}
	}
	return true
}
