# 1043 — Partition Array For Maximum Sum

## Deskripsi

**Soal:** [1043. Partition Array For Maximum Sum](https://leetcode.com/problems/partition-array-for-maximum-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** DP. dp[i] = max sum for prefix ending at i.

## Solusi Go

```go
package main

// LeetCode #1043: Partition Array for Maximum Sum
// https://leetcode.com/problems/partition-array-for-maximum-sum/
// Difficulty: Medium
//
// Approach: DP. dp[i] = max sum for prefix ending at i.
//           For each i, try all partition lengths up to k.
// Time: O(n * k)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3)) // 84
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4)) // 83
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxVal := 0
		for j := 1; j <= k && i-j >= 0; j++ {
			if arr[i-j] > maxVal {
				maxVal = arr[i-j]
			}
			sum := dp[i-j] + maxVal*j
			if sum > dp[i] {
				dp[i] = sum
			}
		}
	}

	return dp[n]
}
```
