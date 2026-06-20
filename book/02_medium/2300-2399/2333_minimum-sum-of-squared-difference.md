# 2333 — Minimum Sum Of Squared Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2333: Minimum Sum of Squared Difference
// https://leetcode.com/problems/minimum-sum-of-squared-difference/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
  // Alokasi slice
	diffs := make([]int, n)
	for i := 0; i < n; i++ {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff = -diff
		}
		diffs[i] = diff
	}

  // Custom sort
	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i] > diffs[j]
	})

	k := k1 + k2
	for i := 0; i < n && k > 0; i++ {
		if diffs[i] == 0 {
			break
		}
		nextDiff := 0
		if i+1 < n {
			nextDiff = diffs[i+1]
		}
		reduce := diffs[i] - nextDiff
		possibleReduce := reduce * (i + 1)
		if possibleReduce <= k {
			k -= possibleReduce
			for j := 0; j <= i; j++ {
				diffs[j] = nextDiff
			}
		} else {
			each := k / (i + 1)
			rem := k % (i + 1)
			for j := 0; j <= i; j++ {
				diffs[j] -= each
				if j < rem {
					diffs[j]--
				}
			}
			k = 0
		}
	}

	var sum int64 = 0
	for _, d := range diffs {
		sum += int64(d) * int64(d)
	}
	return sum
}

func main() {
	// Test case 1
	fmt.Println(minSumSquareDiff([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0))
	// Expected: 579

	// Test case 2
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
	// Expected: 43
}
```
