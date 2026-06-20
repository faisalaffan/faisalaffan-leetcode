# 3727 — Maximum Alternating Sum Of Squares

## Deskripsi

**Soal:** [3727. Maximum Alternating Sum Of Squares](https://leetcode.com/problems/maximum-alternating-sum-of-squares/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maximumAlternatingSumOfSquares(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #3727: Maximum Alternating Sum of Squares
// https://leetcode.com/problems/maximum-alternating-sum-of-squares/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumAlternatingSumOfSquares(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	var ans int64
	for i, v := range nums {
		sq := int64(v) * int64(v)
		if i%2 == 0 {
			ans += sq
		} else {
			ans -= sq
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maximumAlternatingSumOfSquares([]int{1, 2, 3}))
	fmt.Println(maximumAlternatingSumOfSquares([]int{1, -1, 2, -2, 3, -3}))
	fmt.Println(maximumAlternatingSumOfSquares([]int{0, 0, 0}))
}
```
