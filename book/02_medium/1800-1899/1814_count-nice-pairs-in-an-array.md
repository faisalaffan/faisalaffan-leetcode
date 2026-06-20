# 1814 — Count Nice Pairs In An Array

## Deskripsi

**Soal:** [1814. Count Nice Pairs In An Array](https://leetcode.com/problems/count-nice-pairs-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log M) where M = max digit length, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func countNicePairs(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #1814: Count Nice Pairs in an Array
// https://leetcode.com/problems/count-nice-pairs-in-an-array/
// Difficulty: Medium
// Time: O(n log M) where M = max digit length, Space: O(n)

import "fmt"

const mod = 1_000_000_007

func countNicePairs(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[int]int)
	result := 0

	for _, v := range nums {
		key := v - rev(v)
		result = (result + count[key]) % mod
		count[key]++
	}
	return result
}

func rev(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97})) // Expected: 2
	fmt.Println(countNicePairs([]int{13, 10, 35, 24, 76})) // Expected: 4
	fmt.Println(countNicePairs([]int{1, 1, 1, 1})) // Expected: 6
}
```
