package main

// LeetCode #3366: Minimum Array Sum
// https://leetcode.com/problems/minimum-array-sum/
// Difficulty: Medium
// Time: O(n * op1 * op2) Space: O(op1 * op2)

import "fmt"

func main() {
	fmt.Println(minArraySum([]int{2, 8, 3, 19, 3}, 3, 1, 1)) // 23
	fmt.Println(minArraySum([]int{2, 4, 3}, 3, 2, 1))        // 3
}

func minArraySum(nums []int, k int, op1 int, op2 int) int {
	// dp[j1][j2] = min sum using j1 op1 and j2 op2
	dp := make([][]int, op1+1)
	for i := range dp {
		dp[i] = make([]int, op2+1)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0

	for _, x := range nums {
		ndp := make([][]int, op1+1)
		for i := range ndp {
			ndp[i] = make([]int, op2+1)
			for j := range ndp[i] {
				ndp[i][j] = 1 << 60
			}
		}

		for j1 := 0; j1 <= op1; j1++ {
			for j2 := 0; j2 <= op2; j2++ {
				if dp[j1][j2] >= 1<<60 {
					continue
				}
				cur := dp[j1][j2]

				// No operation
				if cur+x < ndp[j1][j2] {
					ndp[j1][j2] = cur + x
				}

				// Only op1 (halve, round up)
				if j1+1 <= op1 {
					val := cur + (x+1)/2
					if val < ndp[j1+1][j2] {
						ndp[j1+1][j2] = val
					}
				}

				// Only op2 (subtract k)
				if j2+1 <= op2 && x >= k {
					val := cur + x - k
					if val < ndp[j1][j2+1] {
						ndp[j1][j2+1] = val
					}
				}

				// Both ops
				if j1+1 <= op1 && j2+1 <= op2 && x >= k {
					// Order 1: subtract then halve
					v1 := (x - k + 1) / 2
					// Order 2: halve then subtract
					half := (x + 1) / 2
					v2 := half - k
					if v2 < 0 {
						v2 = 1 << 60
					}
					best := v1
					if v2 < best {
						best = v2
					}
					val := cur + best
					if val < ndp[j1+1][j2+1] {
						ndp[j1+1][j2+1] = val
					}
				}
			}
		}
		dp = ndp
	}

	ans := 1 << 60
	for j1 := 0; j1 <= op1; j1++ {
		for j2 := 0; j2 <= op2; j2++ {
			if dp[j1][j2] < ans {
				ans = dp[j1][j2]
			}
		}
	}
	return ans
}
