# 1714 — Sum Of Special Evenly Spaced Elements In Array

## Deskripsi

**Soal:** [1714. Sum Of Special Evenly Spaced Elements In Array](https://leetcode.com/problems/sum-of-special-evenly-spaced-elements-in-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Prefix Sum (jumlah kumulatif)

## Solusi Go

```go
package main

// LeetCode #1714: Sum Of Special Evenly-Spaced Elements In Array
// https://leetcode.com/problems/sum-of-special-evenly-spaced-elements-in-array/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	queries := [][]int{{0, 2}, {1, 3}, {2, 1}}
	result := sumOfSpecialEvenlySpacedElements(nums, queries)
	fmt.Printf("Test 1:\n")
	fmt.Printf("Result: %v\nExpected: [25 14 30]\n\n", result)

	nums2 := []int{5, 2, 8, 3, 7, 1, 9, 4, 6}
	queries2 := [][]int{{0, 3}, {2, 2}, {1, 4}}
	result2 := sumOfSpecialEvenlySpacedElements(nums2, queries2)
	fmt.Printf("Test 2:\n")
	fmt.Printf("Result: %v\n", result2)

	nums3 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	queries3 := [][]int{{0, 1}}
	result3 := sumOfSpecialEvenlySpacedElements(nums3, queries3)
	fmt.Printf("\nTest 3 (full sum):\n")
	fmt.Printf("Result: %v\nExpected: [450]\n", result3)
}

func sumOfSpecialEvenlySpacedElements(nums []int, queries [][]int) []int {
	n := len(nums)
	sqrtN := int(math.Sqrt(float64(n)))

	// Precompute prefix sums for small step sizes (y <= sqrtN)
	// preSum[s][i] = sum of nums[i], nums[i-s], nums[i-2s], ... down to index 0
  // Membuat slice 2D untuk DP/tabel
	preSum := make([][]int, sqrtN+1)
	for s := 1; s <= sqrtN; s++ {
		preSum[s] = make([]int, n)
		for i := 0; i < n; i++ {
			if i >= s {
				preSum[s][i] = preSum[s][i-s] + nums[i]
			} else {
				preSum[s][i] = nums[i]
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for idx, q := range queries {
		x, y := q[0], q[1]
		if y <= sqrtN && y > 0 {
			// Use precomputed prefix sums
			last := x + ((n-1-x)/y)*y
			sum := preSum[y][last]
			if x >= y {
				sum -= preSum[y][x-y]
			}
			ans[idx] = sum
		} else if y == 0 {
			ans[idx] = nums[x]
		} else {
			// Large step: brute force
			sum := 0
			for i := x; i < n; i += y {
				sum += nums[i]
			}
			ans[idx] = sum
		}
	}
	return ans
}
```
