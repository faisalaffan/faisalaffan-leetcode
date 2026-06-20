# 2934 — Minimum Operations To Maximize Last Elements In Arrays

## Deskripsi

**Soal:** [2934. Minimum Operations To Maximize Last Elements In Arrays](https://leetcode.com/problems/minimum-operations-to-maximize-last-elements-in-arrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2934: Minimum Operations to Maximize Last Elements in Arrays
// https://leetcode.com/problems/minimum-operations-to-maximize-last-elements-in-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minOperationsArrays([]int{1, 2, 7}, []int{4, 5, 3}))
	fmt.Println(minOperationsArrays([]int{2, 3, 4, 5, 9}, []int{8, 7, 6, 1, 2}))
	fmt.Println(minOperationsArrays([]int{1, 5, 4}, []int{2, 5, 3}))
}

func minOperationsArrays(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := func(x, y int) (cnt int) {
		for i, a := range nums1[:n-1] {
			b := nums2[i]
			if a <= x && b <= y {
				continue
			}
			if !(a <= y && b <= x) {
				return -1
			}
			cnt++
		}
		return
	}
	a, b := f(nums1[n-1], nums2[n-1]), f(nums2[n-1], nums1[n-1])
	if a+b == -2 {
		return -1
	}
	if a < 0 {
		a = 1 << 30
	}
	if b < 0 {
		b = 1 << 30
	}
	if a < b+1 {
		return a
	}
	return b + 1
}
```
