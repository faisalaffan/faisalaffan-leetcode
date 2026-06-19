package main

// LeetCode #69: Sqrt(x)
// https://leetcode.com/problems/sqrtx/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	lo, hi := 1, x/2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(MySqrt(4))
	fmt.Println(MySqrt(8))
	fmt.Println(MySqrt(0))
}
