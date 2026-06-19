package main

// LeetCode #3296: Minimum Number of Seconds to Make Mountain Height Zero
// https://leetcode.com/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
// Difficulty: Medium
// Time: O(n log T) Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minNumberOfSeconds(4, []int{2, 1, 1}))  // 3
	fmt.Println(minNumberOfSeconds(10, []int{3, 2, 2, 4})) // 12
	fmt.Println(minNumberOfSeconds(5, []int{1}))            // 15
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	lo, hi := int64(0), int64(1e18)

	for lo < hi {
		mid := lo + (hi-lo)/2
		if canReduce(mountainHeight, workerTimes, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canReduce(height int, workerTimes []int, t int64) bool {
	var total int64
	for _, wt := range workerTimes {
		// Solve: wt * x * (x+1) / 2 <= t
		// x^2 + x - 2t/wt <= 0
		// x = floor((sqrt(1 + 8*t/wt) - 1) / 2)
		d := math.Sqrt(1.0 + 8.0*float64(t)/float64(wt))
		x := int64((d - 1.0) / 2.0)
		total += x
		if total >= int64(height) {
			return true
		}
	}
	return total >= int64(height)
}
