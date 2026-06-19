package main

// LeetCode #169: Majority Element
// https://leetcode.com/problems/majority-element/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MajorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	fmt.Println(MajorityElement([]int{3, 2, 3}))
	fmt.Println(MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
