# 0238 — Product Of Array Except Self

## Deskripsi

**Soal:** [0238. Product Of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1) excluding output array  
**Kompleksitas Ruang:** O(1) excluding output array

**Algoritma:** —

**Fungsi Solusi:** `func productExceptSelf(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #238: Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/
// Difficulty: Medium
// Time: O(n), Space: O(1) excluding output array

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{0, 0}))
}
```
