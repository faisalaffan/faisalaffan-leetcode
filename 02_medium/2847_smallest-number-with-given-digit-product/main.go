package main

// LeetCode #2847: Smallest Number With Given Digit Product
// https://leetcode.com/problems/smallest-number-with-given-digit-product/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(log n)

import "fmt"

func SmallestNumberWithGivenDigitProduct(n int64) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}

	digits := make([]byte, 0)
	for i := 9; i >= 2; i-- {
		for n%int64(i) == 0 {
			digits = append([]byte{byte('0' + i)}, digits...)
			n /= int64(i)
		}
	}

	if n > 1 {
		return "-1"
	}

	return string(digits)
}

func main() {
	fmt.Println(SmallestNumberWithGivenDigitProduct(36))
	fmt.Println(SmallestNumberWithGivenDigitProduct(17))
	fmt.Println(SmallestNumberWithGivenDigitProduct(1))
}
