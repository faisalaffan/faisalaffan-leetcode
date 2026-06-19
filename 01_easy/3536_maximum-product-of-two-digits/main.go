package main

// LeetCode #3536: Maximum Product of Two Digits
// https://leetcode.com/problems/maximum-product-of-two-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoDigits(34))
	fmt.Println(MaximumProductOfTwoDigits(10))
	fmt.Println(MaximumProductOfTwoDigits(99))
}

// MaximumProductOfTwoDigits returns the maximum product of any two digits in n.
// Time: O(log n). Space: O(1).
func MaximumProductOfTwoDigits(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	maxProd := 0
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			prod := digits[i] * digits[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}
