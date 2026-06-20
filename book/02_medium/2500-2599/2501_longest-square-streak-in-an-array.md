# 2501 — Longest Square Streak In An Array

## Deskripsi

**Soal:** [2501. Longest Square Streak In An Array](https://leetcode.com/problems/longest-square-streak-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2501: Longest Square Streak in an Array
// https://leetcode.com/problems/longest-square-streak-in-an-array/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Sort, use map to track longest streak ending at each value.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2})) // 3 (2 -> 4 -> 16)
	fmt.Println(longestSquareStreak([]int{2, 3, 5, 6, 7}))     // -1
}

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
  // Membuat map untuk pencarian O(1): key → value
	dp := make(map[int]int)
	ans := -1

	for _, v := range nums {
		root := intSqrt(v)
		if root*root == v {
			if prev, ok := dp[root]; ok {
				dp[v] = prev + 1
			} else {
				dp[v] = 1
			}
		} else {
			dp[v] = 1
		}
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	if ans < 2 {
		return -1
	}
	return ans
}

func intSqrt(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid == n {
			return mid
		} else if mid*mid < n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return 0
}
```
