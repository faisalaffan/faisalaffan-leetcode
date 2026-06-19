package main

// LeetCode #326: Power of Three
// https://leetcode.com/problems/power-of-three/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func main() {
	fmt.Println(IsPowerOfThree(27))
	fmt.Println(IsPowerOfThree(0))
	fmt.Println(IsPowerOfThree(-1))
}
