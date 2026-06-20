# 2771 — Longest Non Decreasing Subarray From Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LongestNonDecreasingSubarrayFromTwoArrays(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2771: Longest Non-decreasing Subarray From Two Arrays
// https://leetcode.com/problems/longest-non-decreasing-subarray-from-two-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func LongestNonDecreasingSubarrayFromTwoArrays(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp1, dp2 := 1, 1
	best := 1

	for i := 1; i < n; i++ {
		ndp1, ndp2 := 1, 1
		if nums1[i] >= nums1[i-1] {
			ndp1 = dp1 + 1
		}
		if nums1[i] >= nums2[i-1] {
			if dp2+1 > ndp1 {
				ndp1 = dp2 + 1
			}
		}
		if nums2[i] >= nums1[i-1] {
			ndp2 = dp1 + 1
		}
		if nums2[i] >= nums2[i-1] {
			if dp2+1 > ndp2 {
				ndp2 = dp2 + 1
			}
		}
		dp1, dp2 = ndp1, ndp2
		if dp1 > best {
			best = dp1
		}
		if dp2 > best {
			best = dp2
		}
	}

	return best
}

func main() {
	fmt.Println(LongestNonDecreasingSubarrayFromTwoArrays([]int{1, 3, 2, 1}, []int{2, 2, 3, 4}))
	fmt.Println(LongestNonDecreasingSubarrayFromTwoArrays([]int{1, 2}, []int{3, 1}))
}
```
