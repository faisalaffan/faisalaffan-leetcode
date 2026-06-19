package main

// LeetCode #1056: Confusing Number
// https://leetcode.com/problems/confusing-number/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(confusingNumber(6))   // true (6 -> 9)
	fmt.Println(confusingNumber(89))  // true (89 -> 68)
	fmt.Println(confusingNumber(11))  // false (11 -> 11, same)
	fmt.Println(confusingNumber(25))  // false (invalid digit)
}

// LeetCode submission: confusingNumber
func confusingNumber(n int) bool {
	d := []int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6}
	x, y := n, 0
	for x > 0 {
		v := x % 10
		if d[v] < 0 {
			return false
		}
		y = y*10 + d[v]
		x /= 10
	}
	return y != n
}
