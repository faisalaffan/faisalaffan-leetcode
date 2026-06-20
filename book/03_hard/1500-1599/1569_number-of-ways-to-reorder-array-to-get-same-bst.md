# 1569 — Number Of Ways To Reorder Array To Get Same Bst

## Deskripsi

**Soal:** [1569. Number Of Ways To Reorder Array To Get Same Bst](https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1569: Number of Ways to Reorder Array to Get Same BST
// https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/
// Difficulty: Hard
//
// Combinatorics + recursion:
// 1. The first element is the root.
// 2. Split remaining elements into left (< root) and right (> root).
// 3. The relative order within left and right must be preserved for BST
//    insertion, but we can interleave them arbitrarily.
// 4. Number of ways = C(len(left)+len(right), len(left)) *
//    ways(left) * ways(right) mod M.
// 5. Use Pascal's triangle or precomputed nCr for efficiency.

import (
	"fmt"
)

func main() {
	// Example: [2,1,3] -> 1 (only [2,1,3] works)
	fmt.Println(numOfWays([]int{2, 1, 3}))

	// Additional tests
	fmt.Println(numOfWays([]int{1, 2, 3}))
	fmt.Println(numOfWays([]int{3, 1, 2, 4}))
	fmt.Println(numOfWays([]int{3, 4, 5, 1, 2}))
	fmt.Println(numOfWays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

const MOD = 1_000_000_007

func numOfWays(nums []int) int {
	n := len(nums)
	// Precompute nCr using Pascal's triangle
  // Membuat slice 2D untuk DP/tabel
	comb := make([][]int, n+1)
  // Iterasi seluruh elemen
	for i := range comb {
		comb[i] = make([]int, n+1)
		comb[i][0] = 1
		comb[i][i] = 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % MOD
		}
	}

	// Result minus 1 (exclude original order)
	return (countWays(nums, comb) - 1 + MOD) % MOD
}

func countWays(nums []int, comb [][]int) int {
	if len(nums) <= 1 {
		return 1
	}

	root := nums[0]
  // Membuat slice untuk menyimpan hasil
	left := make([]int, 0)
  // Membuat slice untuk menyimpan hasil
	right := make([]int, 0)

	for i := 1; i < len(nums); i++ {
		if nums[i] < root {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	// C(len(left)+len(right), len(left))
	ways := comb[len(left)+len(right)][len(left)]
	ways = (ways * countWays(left, comb)) % MOD
	ways = (ways * countWays(right, comb)) % MOD

	return ways
}
```
