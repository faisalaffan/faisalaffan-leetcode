# 3654 — Minimum Sum After Divisible Sum Deletions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSumAfterDivisibleSumDeletions(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3654: Minimum Sum After Divisible Sum Deletions
// https://leetcode.com/problems/minimum-sum-after-divisible-sum-deletions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumSumAfterDivisibleSumDeletions(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for i := n/2; i < n; i++ {
		if nums[i]%k == 0 {
			continue
		}
		sum += nums[i]
	}
	for i := 0; i < n/2; i++ {
		if nums[i]%k != 0 {
			sum += nums[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{2, 4, 6, 8}, 2))
	fmt.Println(minimumSumAfterDivisibleSumDeletions([]int{3, 1, 4, 2}, 3))
}
```
