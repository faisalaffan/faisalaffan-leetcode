package main

// LeetCode #306: Additive Number
// https://leetcode.com/problems/additive-number/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	n := len(num)

	for firstEnd := 1; firstEnd <= n/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 1 {
			break
		}
		num1, _ := strconv.ParseInt(num[:firstEnd], 10, 64)

		for secondEnd := firstEnd + 1; max(firstEnd, secondEnd-firstEnd) <= n-secondEnd; secondEnd++ {
			if num[firstEnd] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			num2, _ := strconv.ParseInt(num[firstEnd:secondEnd], 10, 64)

			if isValid(num, num1, num2, secondEnd) {
				return true
			}
		}
	}

	return false
}

func isValid(num string, num1, num2 int64, start int) bool {
	if start == len(num) {
		return false
	}

	for start < len(num) {
		sum := num1 + num2
		sumStr := strconv.FormatInt(sum, 10)

		if !strings.HasPrefix(num[start:], sumStr) {
			return false
		}

		start += len(sumStr)
		num1, num2 = num2, sum
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("12"))
}
