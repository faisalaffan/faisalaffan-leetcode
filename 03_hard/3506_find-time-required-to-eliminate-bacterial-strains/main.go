package main

// LeetCode #3506: Find Time Required to Eliminate Bacterial Strains
// https://leetcode.com/problems/find-time-required-to-eliminate-bacterial-strains/
// Difficulty: Hard [Paid]
//
// Given bacterial strains with initial populations and growth rates, find
// the minimum time required to eliminate all strains using a treatment
// that kills a fixed number of bacteria per unit time.
//
// Approach: Binary search on time. For each time T, check if all strains
// can be eliminated by that time given the kill rate and growth rates.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(findTimeRequired([]int{1, 2, 3}, []int{1, 1, 1}, 2))
	// Example 2
	fmt.Println(findTimeRequired([]int{5, 5, 5}, []int{0, 0, 0}, 3))
	// Edge: single strain
	fmt.Println(findTimeRequired([]int{10}, []int{2}, 3))
	// Edge: no growth
	fmt.Println(findTimeRequired([]int{1, 5}, []int{0, 0}, 1))
}

func findTimeRequired(initial []int, growth []int, killRate int) int {
	n := len(initial)
	if n == 0 {
		return 0
	}

	left, right := 0, math.MaxInt32

	for left < right {
		mid := (left + right) / 2
		if canEliminate(initial, growth, killRate, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canEliminate(initial []int, growth []int, killRate int, time int) bool {
	n := len(initial)
	totalKill := killRate * time

	var totalBacteria int64
	for i := 0; i < n; i++ {
		pop := int64(initial[i]) + int64(growth[i])*int64(time)
		totalBacteria += pop
		if totalBacteria > int64(totalKill) {
			return false
		}
	}

	return totalBacteria <= int64(totalKill)
}
