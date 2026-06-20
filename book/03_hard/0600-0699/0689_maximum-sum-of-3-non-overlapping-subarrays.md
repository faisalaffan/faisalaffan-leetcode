# 0689 — Maximum Sum Of 3 Non Overlapping Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func maxSumOfThreeSubarrays(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #689: Maximum Sum of 3 Non-Overlapping Subarrays
// https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
// Difficulty: Hard
//
// DP: compute best left window position and best right window position,
// then pick the middle window index. O(n) time, O(n) space.

func main() {
	// [1,2,1,2,6,7,5,1], k=2 => [0,3,5]
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
	// [1,2,1,2,1,2,1,2], k=1 => [0,2,4]
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1))
	// [4,5,10,6,11,17,4,11,1,3], k=1 => [0,4,5]
	fmt.Println(maxSumOfThreeSubarrays([]int{4, 5, 10, 6, 11, 17, 4, 11, 1, 3}, 1))
	// [7,13,20,51,12,30,11,15,17,14], k=2 => [2,4,7]
	fmt.Println(maxSumOfThreeSubarrays([]int{7, 13, 20, 51, 12, 30, 11, 15, 17, 14}, 2))
	// Minimal: [1,2,3], k=1 => [0,1,2]
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 3}, 1))
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
  // Alokasi slice
	w := make([]int, n-k+1)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if i >= k {
			sum -= nums[i-k]
		}
		if i >= k-1 {
			w[i-k+1] = sum
		}
	}

  // Alokasi slice
	left := make([]int, len(w))
	best := 0
  // Linear scan O(n)
	for i := 0; i < len(w); i++ {
		if w[i] > w[best] {
			best = i
		}
		left[i] = best
	}

  // Alokasi slice
	right := make([]int, len(w))
	best = len(w) - 1
	for i := len(w) - 1; i >= 0; i-- {
		if w[i] >= w[best] {
			best = i
		}
		right[i] = best
	}

	ans := []int{-1, -1, -1}
	maxSum := 0
	for j := k; j < len(w)-k; j++ {
		i, l := left[j-k], right[j+k]
		total := w[i] + w[j] + w[l]
		if total > maxSum {
			maxSum = total
			ans = []int{i, j, l}
		}
	}
	return ans
}
```
