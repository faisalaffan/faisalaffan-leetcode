package main

// LeetCode #263: Ugly Number
// https://leetcode.com/problems/ugly-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range []int{2, 3, 5} {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

func main() {
	fmt.Println(IsUgly(6))
	fmt.Println(IsUgly(1))
	fmt.Println(IsUgly(14))
}
