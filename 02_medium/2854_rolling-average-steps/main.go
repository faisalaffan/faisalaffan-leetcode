package main

// LeetCode #2854: Rolling Average Steps
// https://leetcode.com/problems/rolling-average-steps/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(k)

import "fmt"

func RollingAverageSteps(steps []int, k int) []float64 {
	n := len(steps)
	if n < k {
		return []float64{}
	}

	result := make([]float64, n-k+1)
	var sum int
	for i := 0; i < k; i++ {
		sum += steps[i]
	}
	result[0] = float64(sum) / float64(k)

	for i := k; i < n; i++ {
		sum += steps[i] - steps[i-k]
		result[i-k+1] = float64(sum) / float64(k)
	}

	return result
}

func main() {
	fmt.Println(RollingAverageSteps([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(RollingAverageSteps([]int{10, 20}, 2))
}
