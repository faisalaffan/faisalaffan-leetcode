package main

// LeetCode #3560: Find Minimum Log Transportation Cost
// https://leetcode.com/problems/find-minimum-log-transportation-cost/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMinimumLogTransportationCost(10, 10, 3))
	fmt.Println(FindMinimumLogTransportationCost(5, 5, 2))
}

// FindMinimumLogTransportationCost returns the minimum transportation cost for logs of size n x m, cutting with factor k.
// Cost = max(0, max(n, m) - k) * k
// Time: O(1). Space: O(1).
func FindMinimumLogTransportationCost(n int, m int, k int) int {
	larger := n
	if m > larger {
		larger = m
	}
	if larger <= k {
		return 0
	}
	return (larger - k) * k
}
