# 1508 — Range Sum Of Sorted Subarray Sums

## Deskripsi

**Soal:** [1508. Range Sum Of Sorted Subarray Sums](https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N^2 log N), Space: O(N^2)  
**Kompleksitas Ruang:** O(N^2)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1508: Range Sum of Sorted Subarray Sums
// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 3, 4))
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 1, 10))
}

func RangeSum(nums []int, n int, left int, right int) int {
	// Time: O(N^2 log N), Space: O(N^2)
	const mod = 1_000_000_007

	// Generate all subarray sums
  // Membuat slice untuk menyimpan hasil
	sums := make([]int, 0, n*(n+1)/2)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			sums = append(sums, sum)
		}
	}

	// Sort
	sort.Ints(sums)

	// Sum from left-1 to right-1
	result := 0
	for i := left - 1; i < right; i++ {
		result = (result + sums[i]) % mod
	}

	return result
}
```
