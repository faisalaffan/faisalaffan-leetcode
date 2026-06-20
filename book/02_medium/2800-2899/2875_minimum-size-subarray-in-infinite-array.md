# 2875 — Minimum Size Subarray In Infinite Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumSizeSubarrayInInfiniteArray(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2875: Minimum Size Subarray in Infinite Array
// https://leetcode.com/problems/minimum-size-subarray-in-infinite-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumSizeSubarrayInInfiniteArray(nums []int, target int) int {
	n := len(nums)
	var totalSum int
	for _, v := range nums {
		totalSum += v
	}

	// If target is 0, we need empty subarray
	if target == 0 {
		return 0
	}

	repeats := target / totalSum
	remainder := target % totalSum

	if remainder == 0 {
		return repeats * n
	}

	// Find minimum subarray with sum == remainder in doubled array
  // Alokasi slice
	extended := make([]int, n*2)
	copy(extended, nums)
	copy(extended[n:], nums)

	best := math.MaxInt32
	left := 0
	sum := 0

	for right := 0; right < len(extended); right++ {
		sum += extended[right]
		for sum > remainder && left <= right {
			sum -= extended[left]
			left++
		}
		if sum == remainder {
			length := right - left + 1
			if length < best {
				best = length
			}
		}
	}

	if best == math.MaxInt32 {
		return -1
	}

	return repeats*n + best
}

func main() {
	fmt.Println(MinimumSizeSubarrayInInfiniteArray([]int{1, 2, 3}, 5))
	fmt.Println(MinimumSizeSubarrayInInfiniteArray([]int{1, 1, 1}, 4))
}
```
