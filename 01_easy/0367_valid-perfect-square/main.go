package main

// LeetCode #367: Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ValidPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq == num {
			return true
		} else if sq < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(ValidPerfectSquare(16))
	fmt.Println(ValidPerfectSquare(14))
	fmt.Println(ValidPerfectSquare(1))
}
