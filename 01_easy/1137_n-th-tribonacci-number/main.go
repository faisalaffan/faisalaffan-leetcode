package main

// LeetCode #1137: N-th Tribonacci Number
// https://leetcode.com/problems/n-th-tribonacci-number/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(tribonacci(4))  // 4
	fmt.Println(tribonacci(25)) // 1389537
	fmt.Println(tribonacci(0))  // 0
}

// LeetCode submission: tribonacci
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
