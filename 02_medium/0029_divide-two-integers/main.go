package main

// LeetCode #29: Divide Two Integers
// https://leetcode.com/problems/divide-two-integers/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend < 0) != (divisor < 0)
	a := abs(dividend)
	b := abs(divisor)
	result := 0

	for a >= b {
		temp := b
		multiple := 1
		for a >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}
		a -= temp
		result += multiple
	}

	if negative {
		return -result
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
	// Test case 1
	fmt.Println(divide(10, 3)) // 3

	// Test case 2
	fmt.Println(divide(7, -3)) // -2

	// Test case 3
	fmt.Println(divide(math.MinInt32, -1)) // 2147483647
}

// Time: O(log n) | Space: O(1)
