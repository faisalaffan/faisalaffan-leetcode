# 3627 — Maximum Median Sum Of Subsequences Of Size 3

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumMedianSumOfSubsequencesOfSizeThree(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3627: Maximum Median Sum of Subsequences of Size 3
// https://leetcode.com/problems/maximum-median-sum-of-subsequences-of-size-3/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{5, 1, 5, 1, 5}))
	// Test case 3
	fmt.Println("Test 3:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{1, 2, 3}))
}

func MaximumMedianSumOfSubsequencesOfSizeThree(nums []int) int {
	n := len(nums)
	maxSum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sub := []int{nums[i], nums[j], nums[k]}
  // Urutkan secara ascending — O(n log n)
				sort.Ints(sub)
				sum := sub[0] + sub[1] + sub[2]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}
```
