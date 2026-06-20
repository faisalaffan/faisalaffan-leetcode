# 1425 — Constrained Subsequence Sum

## Deskripsi

**Soal:** [1425. Constrained Subsequence Sum](https://leetcode.com/problems/constrained-subsequence-sum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func constrainedSubsetSum(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #1425: Constrained Subsequence Sum
// https://leetcode.com/problems/constrained-subsequence-sum/
// Difficulty: Hard

import "fmt"

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n)
	// Monotonic deque storing indices, decreasing dp values
  // Membuat slice untuk menyimpan hasil
	deque := make([]int, 0, n)
	ans := nums[0]

	for i := 0; i < n; i++ {
		// Remove indices out of window
		for len(deque) > 0 && deque[0] < i-k {
			deque = deque[1:]
		}

		dp[i] = nums[i]
		if len(deque) > 0 {
			// max dp in window
			if dp[deque[0]] > 0 {
				dp[i] += dp[deque[0]]
			}
		}

		if dp[i] > ans {
			ans = dp[i]
		}

		// Maintain decreasing deque
		for len(deque) > 0 && dp[deque[len(deque)-1]] <= dp[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	return ans
}

func main() {
	// Example: [10,2,-10,5,20], 2 -> 37
	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
}
```
