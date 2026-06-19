package main

// LeetCode #1017: Convert to Base -2
// https://leetcode.com/problems/convert-to-base-2/
// Difficulty: Medium
//
// Approach: Repeated division by -2. Handle negative remainder.
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(baseNeg2(2))  // "110"
	fmt.Println(baseNeg2(3))  // "111"
	fmt.Println(baseNeg2(4))  // "100"
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n != 0 {
		remainder := n % -2
		n /= -2
		if remainder < 0 {
			remainder += 2
			n++
		}
		result = string(rune('0'+remainder)) + result
	}

	return result
}
