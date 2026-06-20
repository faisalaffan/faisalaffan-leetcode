# 2869 — Minimum Operations To Collect Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumOperationsToCollectElements(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(k)


## 💻 Solusi Go

```go
package main

// LeetCode #2869: Minimum Operations to Collect Elements
// https://leetcode.com/problems/minimum-operations-to-collect-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(k)

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 5))
}

func MinimumOperationsToCollectElements(nums []int, k int) int {
	seen := make([]bool, k+1)
	collected := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= 1 && nums[i] <= k && !seen[nums[i]] {
			seen[nums[i]] = true
			collected++
		}
		if collected == k {
			return len(nums) - i
		}
	}
	return len(nums)
}
```
