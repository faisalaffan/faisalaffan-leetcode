# 2760 — Longest Even Odd Subarray With Threshold

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LongestEvenOddSubarrayWithThreshold(nums []int, threshold int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2760: Longest Even Odd Subarray With Threshold
// https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{3, 2, 5, 4}, 5))
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{4, 5, 2, 1}, 4))
}

func LongestEvenOddSubarrayWithThreshold(nums []int, threshold int) int {
	maxLen := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 || nums[i] > threshold {
			continue
		}
		length := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > threshold {
				break
			}
			if nums[j]%2 == nums[j-1]%2 {
				break
			}
			length++
		}
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
```
