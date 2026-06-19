package main

// LeetCode #1945: Sum of Digits of String After Convert
// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumOfDigitsOfStringAfterConvert("iiii", 1))  // 36
	fmt.Println(SumOfDigitsOfStringAfterConvert("leetcode", 2))  // 6
}

// Time: O(n), Space: O(n)
func SumOfDigitsOfStringAfterConvert(s string, k int) int {
	var digits string
	for i := 0; i < len(s); i++ {
		digits += strconv.Itoa(int(s[i] - 'a' + 1))
	}

	for t := 0; t < k; t++ {
		sum := 0
		for i := 0; i < len(digits); i++ {
			sum += int(digits[i] - '0')
		}
		digits = strconv.Itoa(sum)
	}

	result, _ := strconv.Atoi(digits)
	return result
}
