# 1000 — Minimum Cost To Merge Stones

## Deskripsi

**Soal:** [1000. Minimum Cost To Merge Stones](https://leetcode.com/problems/minimum-cost-to-merge-stones/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Interval DP.

## Solusi Go

```go
package main

// LeetCode #1000: Minimum Cost to Merge Stones
// https://leetcode.com/problems/minimum-cost-to-merge-stones/
// Difficulty: Hard
//
// Approach: Interval DP.
//   dp[i][j] = min cost to merge stones[i:j+1] into (j-i) % (k-1) + 1 piles.
//   We can merge a subarray into 1 pile iff (len-1) % (k-1) == 0.
//   To merge dp[i][j] into 1 pile, we iterate mid where (mid-i) % (k-1) == 0,
//   and dp[i][j] = min(dp[i][mid] + dp[mid+1][j]) + sum(nums[i:j+1]).

import "fmt"

func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2)) // 20
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 3)) // -1
	fmt.Println(mergeStones([]int{1}, 2))           // 0
}

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + stones[i]
	}

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := k; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			dp[i][j] = 1 << 60
			for m := i; m < j; m += k - 1 {
				cost := dp[i][m] + dp[m+1][j]
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
			if (j-i)%(k-1) == 0 {
				dp[i][j] += prefix[j+1] - prefix[i]
			}
		}
	}
	return dp[0][n-1]
}
```
