package main

// LeetCode #3194: Minimum Average of Smallest and Largest Elements
// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/
// Difficulty: Easy

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(MinimumAverageOfSmallestAndLargestElements([]int{1, 9, 8, 3, 10, 5}))
	fmt.Println(MinimumAverageOfSmallestAndLargestElements([]int{1, 2, 3, 7, 8, 9}))
}

// MinimumAverageOfSmallestAndLargestElements returns the minimum average of the smallest and largest elements.
// Time: O(n log n). Space: O(1) (or O(n) due to sorting).
func MinimumAverageOfSmallestAndLargestElements(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	minAvg := math.MaxFloat64
	for i := 0; i < n/2; i++ {
		avg := float64(nums[i]+nums[n-1-i]) / 2.0
		if avg < minAvg {
			minAvg = avg
		}
	}
	return minAvg
}
