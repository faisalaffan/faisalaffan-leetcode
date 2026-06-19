package main

import (
	"fmt"
)

// LeetCode #1227: Airplane Seat Assignment Probability
// https://leetcode.com/problems/airplane-seat-assignment-probability/
// Difficulty: Medium

// n passengers board a plane with n seats. First passenger picks randomly.
// Others: if their seat is free, take it. Otherwise, pick random empty seat.
// Find probability that nth passenger gets their own seat.

// For n=1: 1.0
// For n>=2: 0.5

// Time: O(1)
// Space: O(1)

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1.0
	}
	return 0.5
}

func main() {
	fmt.Printf("%.2f (expected: 1.00)\n", nthPersonGetsNthSeat(1))
	fmt.Printf("%.2f (expected: 0.50)\n", nthPersonGetsNthSeat(2))
	fmt.Printf("%.2f (expected: 0.50)\n", nthPersonGetsNthSeat(100))
}
