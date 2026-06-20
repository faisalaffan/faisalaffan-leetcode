# 3627 — Maximum Median Sum Of Subsequences Of Size 3

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func MaximumMedianSumOfSubsequencesOfSizeThree(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

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
  // Sort O(n log n)
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
