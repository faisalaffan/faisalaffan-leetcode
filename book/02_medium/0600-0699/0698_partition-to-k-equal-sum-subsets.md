# 0698 — Partition To K Equal Sum Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func canPartitionKSubsets(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(k^(n-k) * n!) roughly  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #698: Partition to K Equal Sum Subsets
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
// Difficulty: Medium
// Time: O(k^(n-k) * n!) roughly
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	used := make([]bool, len(nums))

	var backtrack func(start int, currentSum int, subsetsFormed int) bool
	backtrack = func(start int, currentSum int, subsetsFormed int) bool {
		if subsetsFormed == k {
			return true
		}
		if currentSum == target {
			return backtrack(0, 0, subsetsFormed+1)
		}

		for i := start; i < len(nums); i++ {
			if used[i] || currentSum+nums[i] > target {
				continue
			}
			// Pruning: skip duplicate values
			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}

			used[i] = true
			if backtrack(i+1, currentSum+nums[i], subsetsFormed) {
				return true
			}
			used[i] = false

			if currentSum == 0 {
				return false
			}
		}
		return false
	}

	return backtrack(0, 0, 0)
}
```
