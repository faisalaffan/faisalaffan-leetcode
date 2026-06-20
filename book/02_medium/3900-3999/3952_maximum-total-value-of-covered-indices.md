# 3952 — Maximum Total Value Of Covered Indices

## Deskripsi

**Soal:** [3952. Maximum Total Value Of Covered Indices](https://leetcode.com/problems/maximum-total-value-of-covered-indices/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func MaximumTotalValueOfCoveredIndices(nums []int, s string) int64`

> **Ide Kunci:** DP over positions. Token at i can stay (cover i) or move left

## Solusi Go

```go
package main

// LeetCode #3952: Maximum Total Value of Covered Indices
// https://leetcode.com/problems/maximum-total-value-of-covered-indices/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: DP over positions. Token at i can stay (cover i) or move left
// (cover i-1). dp[i][0] = token at i moves left, dp[i][1] = token at i stays.
// No two tokens can cover same index.

import "fmt"

func MaximumTotalValueOfCoveredIndices(nums []int, s string) int64 {
	n := len(nums)
	const negInf int64 = -1 << 60

	dp0, dp1 := int64(0), negInf

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			// No token at i, carry forward previous best
			newDp0 := maxInt64(dp0, dp1)
			dp0, dp1 = newDp0, negInf
		} else {
			// Token at i
			prevDp0, prevDp1 := dp0, dp1

			// Stay at i: covers i, no conflict with any previous decision
			stay := maxInt64(prevDp0, prevDp1) + int64(nums[i])

			// Move left to i-1: covers i-1
			// Requires prev token didn't stay at i-1
			move := negInf
			if i >= 1 {
				// Need prevDp0 (prev token moved left or no prev token)
				move = prevDp0 + int64(nums[i-1])
			}

			dp0 = move
			dp1 = stay
		}
	}

	return maxInt64(dp0, dp1)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(MaximumTotalValueOfCoveredIndices([]int{9, 2, 6, 1}, "0101")) // Expected: 15

	// Example 2
	fmt.Println(MaximumTotalValueOfCoveredIndices([]int{5, 1, 4}, "001")) // Expected: 4

	// Example 3
	fmt.Println(MaximumTotalValueOfCoveredIndices([]int{9, 3, 5}, "011")) // Expected: 14
}
```
