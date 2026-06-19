package main

// LeetCode #1093: Statistics from a Large Sample
// https://leetcode.com/problems/statistics-from-a-large-sample/
// Difficulty: Medium
//
// Approach: Single pass to compute min, max, sum, mode.
//           Two-pointer for median.
// Time: O(n) where n = len(count) = 256
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(sampleStats([]int{0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	// Expected: [1.00000,3.00000,2.37500,2.50000,3.00000]
}

func sampleStats(count []int) []float64 {
	n := 0
	sum := 0
	minVal := -1
	maxVal := 0
	modeVal := 0
	modeCount := 0

	for i, c := range count {
		if c > 0 {
			n += c
			sum += i * c
			if minVal == -1 {
				minVal = i
			}
			maxVal = i
			if c > modeCount {
				modeCount = c
				modeVal = i
			}
		}
	}

	mean := float64(sum) / float64(n)

	// Median
	median := 0.0
	left := n / 2
	right := left + 1
	if n%2 == 1 {
		right = left
	}

	leftVal, rightVal := 0, 0
	acc := 0
	for i, c := range count {
		if c > 0 {
			if acc < left && acc+c >= left {
				leftVal = i
			}
			if acc < right && acc+c >= right {
				rightVal = i
			}
			acc += c
		}
	}
	median = float64(leftVal+rightVal) / 2.0

	return []float64{float64(minVal), float64(maxVal), mean, median, float64(modeVal)}
}
