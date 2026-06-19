package main

// LeetCode #1837: Sum of Digits in Base K
// https://leetcode.com/problems/sum-of-digits-in-base-k/
// Difficulty: Easy

import "fmt"

// Time: O(log_k(n)), Space: O(1)
func SumBase(n int, k int) int {
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}

func main() {
	fmt.Println(SumBase(34, 6))
	fmt.Println(SumBase(10, 10))
	fmt.Println(SumBase(42, 2))
}
