package main

// LeetCode #166: Fraction to Recurring Decimal
// https://leetcode.com/problems/fraction-to-recurring-decimal/
// Difficulty: Medium
// Time: O(n) where n is the length of the repeating cycle, Space: O(n)

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := ""

	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}

	num := abs(numerator)
	den := abs(denominator)

	result += strconv.Itoa(num / den)

	remainder := num % den
	if remainder == 0 {
		return result
	}

	result += "."

	remainderMap := make(map[int]int)
	remainderMap[remainder] = len(result)

	for remainder != 0 {
		remainder *= 10
		result += strconv.Itoa(remainder / den)
		remainder = remainder % den

		if pos, ok := remainderMap[remainder]; ok {
			result = result[:pos] + "(" + result[pos:] + ")"
			break
		}
		remainderMap[remainder] = len(result)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(4, 333))
}
