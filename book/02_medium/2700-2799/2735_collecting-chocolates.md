# 2735 — Collecting Chocolates

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CollectingChocolates(nums []int, x int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2735: Collecting Chocolates
// https://leetcode.com/problems/collecting-chocolates/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func CollectingChocolates(nums []int, x int) int64 {
	n := len(nums)
  // Alokasi slice
	minCost := make([]int, n)
	copy(minCost, nums)

	var best int64
	for i := 0; i < n; i++ {
		best += int64(nums[i])
	}

	for shift := 1; shift < n; shift++ {
		var total int64 = int64(shift) * int64(x)
		for i := 0; i < n; i++ {
			idx := (i + shift) % n
			if nums[idx] < minCost[i] {
				minCost[i] = nums[idx]
			}
			total += int64(minCost[i])
		}
		if total < best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(CollectingChocolates([]int{20, 1, 15}, 5))
	fmt.Println(CollectingChocolates([]int{1, 2, 3}, 4))
}
```
