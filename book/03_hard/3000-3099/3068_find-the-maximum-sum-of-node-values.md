# 3068 — Find The Maximum Sum Of Node Values

## Deskripsi

**Soal:** [3068. Find The Maximum Sum Of Node Values](https://leetcode.com/problems/find-the-maximum-sum-of-node-values/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func maximumValueSum(nums []int, k int, edges [][]int) int64`

> **Ide Kunci:** Sort by gain

## Solusi Go

```go
package main

// LeetCode #3068: Find the Maximum Sum of Node Values
// https://leetcode.com/problems/find-the-maximum-sum-of-node-values/
// Difficulty: Hard
//
// Approach: Sort by gain
// For each node, we can optionally XOR its value with k (gain[i] = (nums[i]^k) - nums[i]).
// Each operation XORs two connected nodes, so the number of XORed nodes must be even.
// Sort gains descending; pair them up; add pair to total if pair sum > 0.
// Tree structure is irrelevant because any even-sized subset is achievable
// via path-based XOR operations.

import (
	"fmt"
	"sort"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	total := int64(0)
  // Membuat slice untuk menyimpan hasil
	gains := make([]int, len(nums))
	for i, v := range nums {
		total += int64(v)
		gains[i] = (v ^ k) - v
	}
	sort.Slice(gains, func(i, j int) bool {
		return gains[i] > gains[j]
	})
	for i := 0; i+1 < len(gains); i += 2 {
		pairSum := gains[i] + gains[i+1]
		if pairSum > 0 {
			total += int64(pairSum)
		}
	}
	return total
}

func main() {
	// Example 1
	fmt.Println("Example 1:", maximumValueSum([]int{1, 2, 1}, 3, [][]int{{0, 1}, {0, 2}}))
	// Expected: 6

	// Example 2
	fmt.Println("Example 2:", maximumValueSum([]int{2, 3}, 7, [][]int{{0, 1}}))
	// Expected: 9

	// Example 3
	fmt.Println("Example 3:", maximumValueSum([]int{7, 8, 9}, 1, [][]int{{0, 1}, {1, 2}}))
	// Expected: 24 (7^1=6, 8^1=9, 9^1=8, all XOR gives decrease except 8^1=9)

	// All zero gains
	fmt.Println("Zero gains:", maximumValueSum([]int{0, 0, 0}, 0, [][]int{{0, 1}, {1, 2}}))
	// Expected: 0

	// Large k value
	fmt.Println("Large k:", maximumValueSum([]int{1, 1, 1}, 100, [][]int{{0, 1}, {1, 2}}))
	// Expected: depends on (1^100) - 1

	// Single edge case
	fmt.Println("Single node:", maximumValueSum([]int{10}, 5, [][]int{}))
	// Expected: 10 (no edges, can't XOR any pair)

	// Two nodes positive gain
	fmt.Println("Two nodes:", maximumValueSum([]int{1, 2}, 3, [][]int{{0, 1}}))
	// 1^3=2 (gain=1), 2^3=1 (gain=-1). Pair sum = 0, not > 0. Total = 1+2 = 3
	// Expected: 3
}
```
