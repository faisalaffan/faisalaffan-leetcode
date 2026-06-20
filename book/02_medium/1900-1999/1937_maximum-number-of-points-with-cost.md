# 1937 — Maximum Number Of Points With Cost

## Deskripsi

**Soal:** [1937. Maximum Number Of Points With Cost](https://leetcode.com/problems/maximum-number-of-points-with-cost/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #1937: Maximum Number of Points with Cost
// https://leetcode.com/problems/maximum-number-of-points-with-cost/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxPoints([][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}))
	fmt.Println(MaxPoints([][]int{{1, 5}, {2, 3}, {4, 2}}))
}

// Time: O(m*n), Space: O(n)
func MaxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
  // Membuat slice untuk menyimpan hasil
	dp := make([]int64, n)
	for j := 0; j < n; j++ {
		dp[j] = int64(points[0][j])
	}

	for i := 1; i < m; i++ {
  // Membuat slice untuk menyimpan hasil
		left := make([]int64, n)
  // Membuat slice untuk menyimpan hasil
		right := make([]int64, n)

		// Left to right: max of dp[k] + k for k <= j
		left[0] = dp[0]
		for j := 1; j < n; j++ {
			if dp[j]+int64(j) > left[j-1] {
				left[j] = dp[j] + int64(j)
			} else {
				left[j] = left[j-1]
			}
		}

		// Right to left: max of dp[k] - k for k >= j
		right[n-1] = dp[n-1] - int64(n-1)
		for j := n - 2; j >= 0; j-- {
			if dp[j]-int64(j) > right[j+1] {
				right[j] = dp[j] - int64(j)
			} else {
				right[j] = right[j+1]
			}
		}

  // Membuat slice untuk menyimpan hasil
		newDp := make([]int64, n)
		for j := 0; j < n; j++ {
			newDp[j] = int64(points[i][j]) + max64(left[j]-int64(j), right[j]+int64(j))
		}
		dp = newDp
	}

	ans := int64(0)
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
```
