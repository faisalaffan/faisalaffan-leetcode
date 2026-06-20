# 3381 — Maximum Subarray Sum With Length Divisible By K

## Deskripsi

**Soal:** [3381. Maximum Subarray Sum With Length Divisible By K](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(k)  
**Kompleksitas Ruang:** O(k)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3381: Maximum Subarray Sum With Length Divisible by K
// https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/
// Difficulty: Medium
// Time: O(n) Space: O(k)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySum([]int{1, 2, 3, 4, 5, 6}, 2)) // 21
	fmt.Println(maxSubarraySum([]int{-1, -2, -3, -4, -5}, 3)) // -6
}

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

  // Membuat slice untuk menyimpan hasil
	minPref := make([]int64, k)
  // Iterasi seluruh elemen
	for i := range minPref {
		minPref[i] = math.MaxInt64
	}
	minPref[0] = 0 // prefix[0] = 0

	var ans int64 = math.MinInt64

	for i := 1; i <= n; i++ {
		r := i % k
		if minPref[r] != math.MaxInt64 {
			val := prefix[i] - minPref[r]
			if val > ans {
				ans = val
			}
		}
		if prefix[i] < minPref[r] {
			minPref[r] = prefix[i]
		}
	}

	if ans == math.MinInt64 {
		return 0
	}
	return ans
}
```
