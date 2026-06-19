package main

import (
	"fmt"
	"math"
)

// LeetCode #1283: Find the Smallest Divisor Given a Threshold
// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/
// Difficulty: Medium

// Binary search the smallest divisor such that sum(ceil(nums[i]/div)) <= threshold.

// Time: O(n log max(nums))
// Space: O(1)

func smallestDivisor(nums []int, threshold int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	lo, hi := 1, maxVal
	for lo < hi {
		mid := lo + (hi-lo)/2
		sum := 0
		for _, v := range nums {
			sum += int(math.Ceil(float64(v) / float64(mid)))
		}
		if sum <= threshold {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Printf("%d (expected: 3)\n", smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Printf("%d (expected: 44)\n", smallestDivisor([]int{44, 22, 33, 11, 1}, 5))
}
