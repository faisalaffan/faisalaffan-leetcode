package main

// LeetCode #3099: Harshad Number
// https://leetcode.com/problems/harshad-number/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfTheDigitsOfHarshadNumber
	fmt.Println(HarshadNumber(18)) // 9
	fmt.Println(HarshadNumber(23)) // -1
}

// Time: O(log n) | Space: O(1)
// LeetCode submission name: sumOfTheDigitsOfHarshadNumber
func HarshadNumber(x int) int {
	sum := 0
	n := x
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
