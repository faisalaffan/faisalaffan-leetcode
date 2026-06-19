package main

// LeetCode #2141: Maximum Running Time of N Computers
// https://leetcode.com/problems/maximum-running-time-of-n-computers/
// Difficulty: Hard
//
// Binary search on the answer. For a given minutes, each battery contributes
// min(capacity, minutes). If sum >= n * minutes, it's feasible.

import "fmt"

func main() {
	fmt.Println(maxRunTime(2, []int{3, 3, 3}))                                    // 4
	fmt.Println(maxRunTime(3, []int{10, 10, 3, 5}))                               // 8
	fmt.Println(maxRunTime(1, []int{1, 2, 3}))                                    // 6
	fmt.Println(maxRunTime(4, []int{10, 10, 10, 10, 5, 5, 5, 5}))                // 15
	fmt.Println(maxRunTime(3, []int{1, 1, 1, 1}))                                 // 1
}

func maxRunTime(n int, batteries []int) int64 {
	canRun := func(minutes int64) bool {
		var total int64
		for _, b := range batteries {
			if int64(b) < minutes {
				total += int64(b)
			} else {
				total += minutes
			}
		}
		return total >= minutes*int64(n)
	}

	lo := int64(0)
	var hi int64
	for _, b := range batteries {
		hi += int64(b)
	}
	hi /= int64(n)

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canRun(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func MaximumRunningTimeOfNComputers() any {
	return maxRunTime(2, []int{3, 3, 3})
}
