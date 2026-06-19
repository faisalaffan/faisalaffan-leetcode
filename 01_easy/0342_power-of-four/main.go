package main

// LeetCode #342: Power of Four
// https://leetcode.com/problems/power-of-four/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func PowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && (n-1)%3 == 0
}

func main() {
	fmt.Println(PowerOfFour(16))
	fmt.Println(PowerOfFour(5))
	fmt.Println(PowerOfFour(1))
}
