# 3654 — Minimum Sum After Divisible Sum Deletions

## Deskripsi

**Soal:** [3654. Minimum Sum After Divisible Sum Deletions](https://leetcode.com/problems/minimum-sum-after-divisible-sum-deletions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumSumAfterDivisibleSumDeletions(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #3654: Minimum Sum After Divisible Sum Deletions
// https://leetcode.com/problems/minimum-sum-after-divisible-sum-deletions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumSumAfterDivisibleSumDeletions(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for i := n/2; i < n; i++ {
		if nums[i]%k == 0 {
			continue
		}
		sum += nums[i]
	}
	for i := 0; i < n/2; i++ {
		if nums[i]%k != 0 {
			sum += nums[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{2, 4, 6, 8}, 2))
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{3, 1, 4, 2}, 3))
}
```
