package main

// LeetCode #3119: Maximum Number of Potholes That Can Be Fixed
// https://leetcode.com/problems/maximum-number-of-potholes-that-can-be-fixed/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxPotholes(road string, budget int) int {
	var segs []int
	count := 0
	for i := 0; i < len(road); i++ {
		if road[i] == 'x' {
			count++
		} else {
			if count > 0 {
				segs = append(segs, count)
				count = 0
			}
		}
	}
	if count > 0 {
		segs = append(segs, count)
	}

	sort.Slice(segs, func(i, j int) bool {
		return segs[i] > segs[j]
	})

	ans := 0
	for _, seg := range segs {
		cost := seg + 1
		if budget >= cost {
			budget -= cost
			ans += seg
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPotholes("...xxx..xx", 7))  // Expected: 5
	fmt.Println(maxPotholes("..xxxxx", 4))     // Expected: 3
	fmt.Println(maxPotholes("x.x.x.x", 10))    // Expected: 4
}
