package main

// LeetCode #2769: Find the Maximum Achievable Number
// https://leetcode.com/problems/find-the-maximum-achievable-number/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumAchievableNumber(4, 1))
	fmt.Println(FindTheMaximumAchievableNumber(3, 2))
}

func FindTheMaximumAchievableNumber(num int, t int) int {
	return num + 2*t
}
