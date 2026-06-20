# 1630 — Arithmetic Subarrays

## Deskripsi

**Soal:** [1630. Arithmetic Subarrays](https://leetcode.com/problems/arithmetic-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(M * N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1630: Arithmetic Subarrays
// https://leetcode.com/problems/arithmetic-subarrays/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CheckArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(CheckArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
}

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	// Time: O(M * N log N), Space: O(N)
  // Membuat slice untuk menyimpan hasil
	result := make([]bool, len(l))

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(l); i++ {
  // Membuat slice untuk menyimpan hasil
		sub := make([]int, r[i]-l[i]+1)
		copy(sub, nums[l[i]:r[i]+1])
		result[i] = isArithmetic(sub)
	}

	return result
}

func isArithmetic(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}

	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
```
