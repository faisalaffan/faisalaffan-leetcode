# 2547 — Minimum Cost To Split An Array

## Deskripsi

**Soal:** [2547. Minimum Cost To Split An Array](https://leetcode.com/problems/minimum-cost-to-split-an-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minCost(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2547: Minimum Cost to Split an Array
// https://leetcode.com/problems/minimum-cost-to-split-an-array/
// Difficulty: Hard

import "fmt"

// minCost uses DP where dp[i] = min cost for prefix nums[0..i-1].
// For each i, we try all j < i as the start of the last subarray,
// tracking the "trimmed" count: number of distinct values whose
// frequency in the subarray >= 2. Cost = trimmed + k.
//
// Complexity: O(n^2) time, O(n) space
func minCost(nums []int, k int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = 1 << 60
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
  // Membuat map untuk pencarian O(1): key → value
		freq := make(map[int]int)
		trimmed := 0
		for j := i - 1; j >= 0; j-- {
			x := nums[j]
			freq[x]++
			if freq[x] == 2 {
				// This value now appears more than once for the first time
				trimmed++
			}
			cost := trimmed + k
			if dp[j]+cost < dp[i] {
				dp[i] = dp[j] + cost
			}
		}
	}
	return dp[n]
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [1,2,1,2,1], k=2 ->", minCost([]int{1, 2, 1, 2, 1}, 2)) // 4

	// Additional test cases
	fmt.Println("Test 2: [1,2,1,2,1], k=5 ->", minCost([]int{1, 2, 1, 2, 1}, 5)) // 7 (no split better than whole)
	fmt.Println("Test 3: [1,2,3,4,5], k=2 ->", minCost([]int{1, 2, 3, 4, 5}, 2)) // all unique -> each singly
	fmt.Println("Test 4: [0,0,0,0], k=1 ->", minCost([]int{0, 0, 0, 0}, 1))      // all same: trimmed=1 per subarray
	fmt.Println("Test 5: [] ->", minCost([]int{}, 5))                             // 0
}
```
