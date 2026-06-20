# 2348 — Number Of Zero Filled Subarrays

## Deskripsi

**Soal:** [2348. Number Of Zero Filled Subarrays](https://leetcode.com/problems/number-of-zero-filled-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func zeroFilledSubarray(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #2348: Number of Zero-Filled Subarrays
// https://leetcode.com/problems/number-of-zero-filled-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	var total int64 = 0
	var count int64 = 0

	for _, v := range nums {
		if v == 0 {
			count++
			total += count
		} else {
			count = 0
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	// Expected: 6

	// Test case 2
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
	// Expected: 9

	// Test case 3
	fmt.Println(zeroFilledSubarray([]int{0}))
	// Expected: 1
}
```
