package main

// LeetCode #2283: Check if Number Has Equal Digit Count and Digit Value
// https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("1210")) // true
	fmt.Println(CheckIfNumberHasEqualDigitCountAndDigitValue("030"))  // false
}

func CheckIfNumberHasEqualDigitCountAndDigitValue(num string) bool {
	count := [10]int{}
	for i := 0; i < len(num); i++ {
		count[num[i]-'0']++
	}
	for i := 0; i < len(num); i++ {
		if count[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
