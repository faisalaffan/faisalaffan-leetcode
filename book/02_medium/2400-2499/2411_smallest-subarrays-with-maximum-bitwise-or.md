# 2411 — Smallest Subarrays With Maximum Bitwise Or

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func smallestSubarrays(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * 30)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2411: Smallest Subarrays With Maximum Bitwise OR
// https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/
// Difficulty: Medium
// Time: O(n * 30) | Space: O(1)
// For each position i, find min length subarray starting at i with max possible OR.
// Track for each bit, the nearest position to the right where it's set.

import "fmt"

func main() {
	fmt.Println(smallestSubarrays([]int{1, 0, 2, 1, 3})) // [3, 3, 2, 2, 1]
	fmt.Println(smallestSubarrays([]int{1, 2}))          // [2, 1]
}

func smallestSubarrays(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	ans := make([]int, n)
  // Alokasi slice
	last := make([]int, 30) // for each bit, last position where it's set

	bitPos := -1
	for i := n - 1; i >= 0; i-- {
		maxDist := 1
		for b := 0; b < 30; b++ {
			if nums[i]>>b&1 == 1 {
				last[b] = i
			}
			bitPos = last[b]
			if bitPos != 0 {
				dist := bitPos - i + 1
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
		ans[i] = maxDist
	}
	return ans
}
```
