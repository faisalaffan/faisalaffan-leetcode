# 0891 — Sum Of Subsequence Widths

## Deskripsi

**Soal:** [0891. Sum Of Subsequence Widths](https://leetcode.com/problems/sum-of-subsequence-widths/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func sumSubseqWidths(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #891: Sum of Subsequence Widths
// https://leetcode.com/problems/sum-of-subsequence-widths/
// Difficulty: Hard
//
// Sort the array. For a sorted array, each element nums[i] is:
//   - The maximum of 2^i subsequences (every subsequence of the first i+1 elements
//     that includes nums[i])
//   - The minimum of 2^(n-1-i) subsequences (every subsequence of the last n-i elements
//     that includes nums[i])
//
// Contribution: nums[i] * (2^i - 2^(n-1-i))
// Sum all contributions mod 1e9+7.

import (
	"fmt"
	"sort"
)

const mod891 = 1_000_000_007

func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	// Precompute powers of 2.
  // Membuat slice untuk menyimpan hasil
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = (pow2[i-1] * 2) % mod891
	}

	result := 0
	for i := 0; i < n; i++ {
		contribution := (pow2[i] - pow2[n-1-i] + mod891) % mod891
		result = (result + (nums[i] % mod891)*contribution) % mod891
	}
	return result
}

func main() {
	// Example 1: [2,1,3] -> 6
	// Sorted: [1,2,3]; widths: |1-?| subsequences
	// 1: 2^0-2^2 = 1-4 = -3; 2: 2^1-2^1 = 0; 3: 2^2-2^0 = 4-1 = 3
	// 1*(-3) + 2*0 + 3*3 = -3+0+9 = 6
	fmt.Println("Test 1:", sumSubseqWidths([]int{2, 1, 3})) // 6

	// Example 2: [1] -> 0 (only one element, no subsequence of length >= 2)
	fmt.Println("Test 2:", sumSubseqWidths([]int{1})) // 0

	// Example 3: [1,2] -> 1
	// Sorted: [1,2]; 1: 1-2=-1, 2: 2-1=1, result = -1+2 = 1
	fmt.Println("Test 3:", sumSubseqWidths([]int{1, 2})) // 1

	// Larger: [5,2,1,4] -> ?
	fmt.Println("Test 4:", sumSubseqWidths([]int{5, 2, 1, 4}))

	// All same: [7,7,7] -> 0 (width is always 0)
	fmt.Println("Test 5:", sumSubseqWidths([]int{7, 7, 7})) // 0
}
```
