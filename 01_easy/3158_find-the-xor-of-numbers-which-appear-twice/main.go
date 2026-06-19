package main

// LeetCode #3158: Find the XOR of Numbers Which Appear Twice
// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: duplicateNumbersXOR
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 2, 1})) // 3
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 3}))    // 0
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: duplicateNumbersXOR
func FindTheXorOfNumbersWhichAppearTwice(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	xor := 0
	for v, f := range freq {
		if f == 2 {
			xor ^= v
		}
	}
	return xor
}
