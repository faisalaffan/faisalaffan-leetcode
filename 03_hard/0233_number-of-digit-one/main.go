package main

// LeetCode #233: Number of Digit One
// https://leetcode.com/problems/number-of-digit-one/
// Difficulty: Hard

import "fmt"

func countDigitOne(n int) int {
	count := 0
	if n <= 0 {
		return 0
	}

	factor := 1
	for factor <= n {
		lower := n % factor
		cur := (n / factor) % 10
		higher := n / (factor * 10)

		switch cur {
		case 0:
			count += higher * factor
		case 1:
			count += higher*factor + lower + 1
		default:
			count += (higher + 1) * factor
		}

		// Check overflow
		if factor > n/10 {
			break
		}
		factor *= 10
	}

	return count
}

func main() {
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(0))
}
