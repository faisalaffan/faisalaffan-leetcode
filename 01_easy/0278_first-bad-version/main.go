package main

// LeetCode #278: First Bad Version
// https://leetcode.com/problems/first-bad-version/
// Difficulty: Easy

import "fmt"

var firstBad int

func isBadVersion(version int) bool {
	return version >= firstBad
}

// Time: O(log n) | Space: O(1)
func FirstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	firstBad = 4
	fmt.Println(FirstBadVersion(5))
	firstBad = 1
	fmt.Println(FirstBadVersion(1))
}
