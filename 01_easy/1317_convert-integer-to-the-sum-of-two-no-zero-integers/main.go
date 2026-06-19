package main

// LeetCode #1317: Convert Integer to the Sum of Two No-Zero Integers
// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
// Difficulty: Easy
//
// LeetCode submission: func getNoZeroIntegers(n int) []int

import "fmt"

func main() {
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(2))    // [1 1]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(11))   // [2 9]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(1010)) // [122 888]
}

// Time: O(n log n), Space: O(1)
func ConvertIntegerToTheSumOfTwoNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return []int{}
}

func hasZero(x int) bool {
	for x > 0 {
		if x%10 == 0 {
			return true
		}
		x /= 10
	}
	return false
}
