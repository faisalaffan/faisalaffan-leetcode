# 1818 — Minimum Absolute Sum Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1818: Minimum Absolute Sum Difference
// https://leetcode.com/problems/minimum-absolute-sum-difference/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

const mod = 1_000_000_007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, nums1)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)

	total := 0
	maxReduction := 0

	for i := 0; i < n; i++ {
		origDiff := abs(nums1[i] - nums2[i])
		total = (total + origDiff) % mod

		// Find closest value to nums2[i] in sorted nums1
		idx := sort.SearchInts(sorted, nums2[i])
		if idx < n {
			maxReduction = max(maxReduction, origDiff-abs(sorted[idx]-nums2[i]))
		}
		if idx > 0 {
			maxReduction = max(maxReduction, origDiff-abs(sorted[idx-1]-nums2[i]))
		}
	}

	return (total - maxReduction + mod) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5})) // Expected: 3
	fmt.Println(minAbsoluteSumDiff([]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10})) // Expected: 0
	fmt.Println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4})) // Expected: 20
}
```
