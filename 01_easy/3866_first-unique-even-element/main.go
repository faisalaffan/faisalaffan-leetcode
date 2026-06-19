package main

// LeetCode #3866: First Unique Even Element
// https://leetcode.com/problems/first-unique-even-element/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FirstUniqueEvenElement([]int{3, 4, 2, 5, 4, 6}))
	fmt.Println(FirstUniqueEvenElement([]int{4, 4}))
	fmt.Println(FirstUniqueEvenElement([]int{1, 3, 5}))
}

// Time: O(n)
// Space: O(n)
func FirstUniqueEvenElement(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, v := range nums {
		if v%2 == 0 && freq[v] == 1 {
			return v
		}
	}
	return -1
}
