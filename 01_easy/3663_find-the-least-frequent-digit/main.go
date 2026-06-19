package main

// LeetCode #3663: Find The Least Frequent Digit
// https://leetcode.com/problems/find-the-least-frequent-digit/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheLeastFrequentDigit(1553322))
	fmt.Println(FindTheLeastFrequentDigit(723344511))
}

// Time: O(log n) - number of digits
// Space: O(1)
func FindTheLeastFrequentDigit(n int) int {
	cnt := [10]int{}
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}

	minCnt := math.MaxInt
	ans := 0
	for d := 0; d <= 9; d++ {
		if cnt[d] > 0 && (cnt[d] < minCnt || (cnt[d] == minCnt && d < ans)) {
			minCnt = cnt[d]
			ans = d
		}
	}
	return ans
}
