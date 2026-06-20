# 3356 — Zero Array Transformation Ii

## Deskripsi

**Soal:** [3356. Zero Array Transformation Ii](https://leetcode.com/problems/zero-array-transformation-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((n + q) log q) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3356: Zero Array Transformation II
// https://leetcode.com/problems/zero-array-transformation-ii/
// Difficulty: Medium
// Time: O((n + q) log q) Space: O(n)

import "fmt"

func main() {
	fmt.Println(minZeroArray([]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}})) // 2
	fmt.Println(minZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}))         // -1
}

func minZeroArray(nums []int, queries [][]int) int {
	m := len(queries)

	// Check if already zero
	allZero := true
	for _, v := range nums {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	// Check if impossible
	if !canMakeZero(nums, queries, m) {
		return -1
	}

	lo, hi := 1, m
	for lo < hi {
		mid := (lo + hi) / 2
		if canMakeZero(nums, queries, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canMakeZero(nums []int, queries [][]int, k int) bool {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n+1)
	for i := 0; i < k; i++ {
		l, r, val := queries[i][0], queries[i][1], queries[i][2]
		diff[l] += val
		diff[r+1] -= val
	}
	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```
