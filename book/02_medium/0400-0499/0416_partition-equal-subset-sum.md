# 0416 — Partition Equal Subset Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func canPartition(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * sum)  |  **Ruang:** O(sum)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #416: Partition Equal Subset Sum
// https://leetcode.com/problems/partition-equal-subset-sum/
// Difficulty: Medium
// Time: O(n * sum) | Space: O(sum)

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for s := target; s >= num; s-- {
			if dp[s-num] {
				dp[s] = true
			}
		}
	}
	return dp[target]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canPartition([]int{1, 5, 11, 5}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canPartition([]int{1, 2, 3, 5}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", canPartition([]int{2, 2, 2, 2}))
	// Expected: true
}
```
