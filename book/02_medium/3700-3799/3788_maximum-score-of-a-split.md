# 3788 — Maximum Score Of A Split

## Deskripsi

**Soal:** [3788. Maximum Score Of A Split](https://leetcode.com/problems/maximum-score-of-a-split/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maximumScoreOfASplit(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #3788: Maximum Score of a Split
// https://leetcode.com/problems/maximum-score-of-a-split/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func maximumScoreOfASplit(nums []int) int64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	suf := make([]int64, n)
	suf[n-1] = int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		suf[i] = int64(nums[i])
		if suf[i+1] < suf[i] {
			suf[i] = suf[i+1]
		}
	}

	var pre int64
	var ans int64 = math.MinInt64
	for i := 0; i < n-1; i++ {
		pre += int64(nums[i])
		score := pre - suf[i+1]
		if score > ans {
			ans = score
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumScoreOfASplit([]int{10, -1, 3, -4, -5}))
	fmt.Println(maximumScoreOfASplit([]int{1, 2, 3, 4}))
	fmt.Println(maximumScoreOfASplit([]int{-5, -3, -1}))
}
```
