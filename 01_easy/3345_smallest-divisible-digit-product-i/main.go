package main

// LeetCode #3345: Smallest Divisible Digit Product I
// https://leetcode.com/problems/smallest-divisible-digit-product-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestDivisibleDigitProductI(10, 2))
	fmt.Println(SmallestDivisibleDigitProductI(15, 3))
}

// digitProduct returns the product of digits of n.
func digitProduct(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n /= 10
	}
	return product
}

// SmallestDivisibleDigitProductI returns the smallest number >= n whose digit product is divisible by t.
// Time: O(answer * log n). Space: O(1).
func SmallestDivisibleDigitProductI(n int, t int) int {
	for {
		dp := digitProduct(n)
		if dp%t == 0 {
			return n
		}
		n++
	}
}
