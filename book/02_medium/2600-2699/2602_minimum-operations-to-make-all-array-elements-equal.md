# 2602 — Minimum Operations To Make All Array Elements Equal

## Deskripsi

**Soal:** [2602. Minimum Operations To Make All Array Elements Equal](https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((n+q) log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minOperations(nums []int, queries []int) []int64`

## Solusi Go

```go
package main

// LeetCode #2602: Minimum Operations to Make All Array Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/
// Difficulty: Medium
// Time: O((n+q) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)

  // Membuat slice untuk menyimpan hasil
	prefix := make([]int64, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + int64(v)
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int64, len(queries))
	for i, q := range queries {
		idx := sort.SearchInts(nums, q)
		leftCount := int64(idx)
		rightCount := int64(n - idx)
		leftSum := prefix[idx]
		rightSum := prefix[n] - prefix[idx]
		ops := int64(q)*leftCount - leftSum + rightSum - int64(q)*rightCount
		ans[i] = ops
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{3, 1, 6, 8}, []int{1, 5}))
	// Expected: [8, 10]? Let me check...

	// Test case 2
	fmt.Println("Test 2:", minOperations([]int{2, 4, 6, 8}, []int{4, 5}))
	// Expected: [4, 4]

	// Test case 3
	fmt.Println("Test 3:", minOperations([]int{1, 2, 3, 4, 5}, []int{3}))
	// Expected: [6]
}
```
