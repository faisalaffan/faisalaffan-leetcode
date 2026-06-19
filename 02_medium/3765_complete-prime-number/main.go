package main

// LeetCode #3765: Complete Prime Number
// https://leetcode.com/problems/complete-prime-number/
// Difficulty: Medium
// Time: O(d * sqrt(n)) | Space: O(1)

import "fmt"

func completePrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}

	suffix := 0
	power := 1
	x := num

	for x > 0 {
		// Build suffix: prepend last digit
		suffix = power*(x%10) + suffix
		power *= 10

		if !isPrime(suffix) {
			return false
		}
		if !isPrime(x) {
			return false
		}
		x /= 10
	}
	return true
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(completePrimeNumber(23))
	fmt.Println(completePrimeNumber(39))
	fmt.Println(completePrimeNumber(7))
}
