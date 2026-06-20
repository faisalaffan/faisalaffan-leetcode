# 1027 — Longest Arithmetic Subsequence

## Deskripsi

**Soal:** [1027. Longest Arithmetic Subsequence](https://leetcode.com/problems/longest-arithmetic-subsequence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** HashMap (tabel pencarian O(1)), Dynamic Programming (DP)

> **Ide Kunci:** DP with hash map per index tracking difference -> length

## Solusi Go

```go
package main

// LeetCode #1027: Longest Arithmetic Subsequence
// https://leetcode.com/problems/longest-arithmetic-subsequence/
// Difficulty: Medium
//
// Approach: DP with hash map per index tracking difference -> length
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))    // 4
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10})) // 3
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8})) // 4
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

  // Membuat slice untuk menyimpan hasil
	dp := make([]map[int]int, n)
	result := 2

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			length := 2
			if prev, ok := dp[j][diff]; ok {
				length = prev + 1
			}
			dp[i][diff] = length
			if length > result {
				result = length
			}
		}
	}

	return result
}
```
