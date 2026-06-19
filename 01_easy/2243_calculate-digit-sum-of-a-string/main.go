package main

// LeetCode #2243: Calculate Digit Sum of a String
// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CalculateDigitSumOfAString("11111222223", 3)) // "135"
	fmt.Println(CalculateDigitSumOfAString("00000000", 3))    // "000"
}

// Time: O(n), Space: O(n)
func CalculateDigitSumOfAString(s string, k int) string {
	for len(s) > k {
		var next string
		for i := 0; i < len(s); i += k {
			end := i + k
			if end > len(s) {
				end = len(s)
			}
			sum := 0
			for j := i; j < end; j++ {
				sum += int(s[j] - '0')
			}
			next += strconv.Itoa(sum)
		}
		s = next
	}
	return s
}
