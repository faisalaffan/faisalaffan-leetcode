package main

// LeetCode #2604: Minimum Time to Eat All Grains
// https://leetcode.com/problems/minimum-time-to-eat-all-grains/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// minimumTime finds minimum time for all hens to eat all grains.
// Hens move speed 1 on integer line simultaneously and independently.
// Each hen can eat multiple grains; eating takes negligible time.
//
// Binary search answer T. For a given T, greedily assign grains to hens
// from left to right. For a hen at h with leftmost uneaten grain g < h,
// there are two strategies:
//   A) Go left first to g, then right: maxReach = T + 2*g - h
//   B) Go right first, then left to g:   maxReach = (T + h + g) / 2
// The hen uses the better strategy (max of both).
//
// Complexity: O((n+m) log maxPos) time, O(1) space
func minimumTime(hens []int, grains []int) int {
	sort.Ints(hens)
	sort.Ints(grains)

	canEat := func(T int) bool {
		gIdx := 0
		m := len(grains)
		for _, h := range hens {
			if gIdx >= m {
				break
			}
			g := grains[gIdx]
			if g < h {
				// Leftmost grain is to the left of hen
				dist := h - g
				if dist > T {
					return false
				}
				// Strategy A: go left first, then right
				maxA := T + 2*g - h
				// Strategy B: go right first, then left to cover g
				maxB := (T + h + g) / 2
				maxPos := maxA
				if maxB > maxPos {
					maxPos = maxB
				}
				if maxPos < h {
					maxPos = h
				}
				for gIdx < m && grains[gIdx] <= maxPos {
					gIdx++
				}
			} else {
				// All remaining grains are at or right of hen
				maxPos := h + T
				for gIdx < m && grains[gIdx] <= maxPos {
					gIdx++
				}
			}
		}
		return gIdx >= m
	}

	lo, hi := 0, 2_000_000_000
	for lo < hi {
		mid := lo + (hi-lo)/2
		if canEat(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: hens=[3,6,7], grains=[2,4,6,8] ->",
		minimumTime([]int{3, 6, 7}, []int{2, 4, 6, 8}))

	// Official example
	fmt.Println("Test 2: hens=[3,6,7], grains=[2,4,7,9] ->",
		minimumTime([]int{3, 6, 7}, []int{2, 4, 7, 9}))

	// Additional test cases
	fmt.Println("Test 3: hens=[1,10], grains=[5] ->",
		minimumTime([]int{1, 10}, []int{5}))

	fmt.Println("Test 4: hens=[0,4], grains=[2] ->",
		minimumTime([]int{0, 4}, []int{2}))

	fmt.Println("Test 5: hens=[0], grains=[0] ->",
		minimumTime([]int{0}, []int{0}))
}
