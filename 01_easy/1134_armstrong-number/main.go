package main

// LeetCode #1134: Armstrong Number
// https://leetcode.com/problems/armstrong-number/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isArmstrong(153))  // true
	fmt.Println(isArmstrong(123))  // false
	fmt.Println(isArmstrong(1))    // true
}

// LeetCode submission: isArmstrong
func isArmstrong(n int) bool {
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		p := 1
		for i := 0; i < digits; i++ {
			p *= d
		}
		sum += p
	}
	return sum == n
}
