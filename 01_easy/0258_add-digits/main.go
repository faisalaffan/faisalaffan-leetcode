package main

// LeetCode #258: Add Digits
// https://leetcode.com/problems/add-digits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func AddDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	fmt.Println(AddDigits(38))
	fmt.Println(AddDigits(0))
	fmt.Println(AddDigits(9))
}
