# 3724 — Minimum Operations To Transform Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumOperationsToTransformArray(nums1 []int, nums2 []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3724: Minimum Operations to Transform Array
// https://leetcode.com/problems/minimum-operations-to-transform-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformArray(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	last := nums2[n]
	var ops int64 = 1
	var extra int64 = 1 << 60

	for i := 0; i < n; i++ {
		lo, hi := nums1[i], nums2[i]
		if lo > hi {
			lo, hi = hi, lo
		}
		ops += int64(hi - lo)

		if lo <= last && last <= hi {
			extra = 0
		} else if last < lo {
			if int64(1+lo-last) < extra {
				extra = int64(1 + lo - last)
			}
		} else {
			if int64(1+last-hi) < extra {
				extra = int64(1 + last - hi)
			}
		}
	}

	return ops + extra
}

func main() {
	fmt.Println(minimumOperationsToTransformArray([]int{2, 8}, []int{1, 7, 3}))
	fmt.Println(minimumOperationsToTransformArray([]int{1, 2}, []int{3, 4, 5}))
	fmt.Println(minimumOperationsToTransformArray([]int{5, 5}, []int{5, 5, 5}))
}
```
