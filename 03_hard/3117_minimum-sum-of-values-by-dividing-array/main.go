package main

// LeetCode #3117: Minimum Sum of Values by Dividing Array
// https://leetcode.com/problems/minimum-sum-of-values-by-dividing-array/
// Difficulty: Hard
//
// Partition nums into m contiguous subarrays such that the AND of the i-th
// subarray equals andValues[i]. Minimize the sum of the last elements of each
// subarray. Use DP with maps tracking (completed_segments, current_AND) -> min_sum.

import (
	"fmt"
)

const ALL_ONES = (1 << 20) - 1

func main() {
	// Example: [1,4,3,3,2], [0,3,3,2] -> 12
	nums := []int{1, 4, 3, 3, 2}
	andValues := []int{0, 3, 3, 2}
	fmt.Println(minimumSumOfValuesByDividingArray(nums, andValues))
}

func minimumSumOfValuesByDividingArray(nums []int, andValues []int) int {
	m := len(andValues)
	dp := make([]map[int]int, m+1)
	for j := 0; j <= m; j++ {
		dp[j] = make(map[int]int)
	}
	dp[0][ALL_ONES] = 0

	for _, x := range nums {
		ndp := make([]map[int]int, m+1)
		for j := 0; j <= m; j++ {
			ndp[j] = make(map[int]int)
		}

		for j := 0; j <= m; j++ {
			for andVal, sum := range dp[j] {
				if andVal == ALL_ONES {
					// Start new segment at x, don't close
					if val, ok := ndp[j][x]; !ok || sum < val {
						ndp[j][x] = sum
					}
					// Start new segment at x and immediately close (single element segment)
					if j < m && x == andValues[j] {
						if val, ok := ndp[j+1][ALL_ONES]; !ok || sum+x < val {
							ndp[j+1][ALL_ONES] = sum + x
						}
					}
				} else {
					// Extend segment with x
					newAnd := andVal & x
					// Extend, don't close
					if val, ok := ndp[j][newAnd]; !ok || sum < val {
						ndp[j][newAnd] = sum
					}
					// Extend and close segment at x
					if j < m && newAnd == andValues[j] {
						if val, ok := ndp[j+1][ALL_ONES]; !ok || sum+x < val {
							ndp[j+1][ALL_ONES] = sum + x
						}
					}
				}
			}
		}

		dp = ndp
	}

	if ans, ok := dp[m][ALL_ONES]; ok {
		return ans
	}
	return -1
}
