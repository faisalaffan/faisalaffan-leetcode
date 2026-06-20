# 1300 — Sum Of Mutated Array Closest To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findBestValue(arr []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** O(n log n + n log max(arr))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1300: Sum of Mutated Array Closest to Target
// https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/
// Difficulty: Medium

// Find integer value such that sum(arr[i] if arr[i] < value else value)
// is as close to target as possible. If tie, return smaller value.

// Time: O(n log n + n log max(arr))
// Space: O(1)

func findBestValue(arr []int, target int) int {
  // Sort O(n log n)
	sort.Ints(arr)
	n := len(arr)

  // Alokasi slice
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	lo, hi := 0, arr[n-1]
	bestVal, minDiff := 0, target

	for lo <= hi {
		mid := lo + (hi-lo)/2

		// Find first index > mid
		idx := sort.Search(n, func(i int) bool {
			return arr[i] > mid
		})

		sum := prefix[idx] + mid*(n-idx)
		diff := abs(sum - target)

		if diff < minDiff || (diff == minDiff && mid < bestVal) {
			bestVal = mid
			minDiff = diff
		}

		if sum < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return bestVal
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%d (expected: 3)\n", findBestValue([]int{4, 9, 3}, 10))
	fmt.Printf("%d (expected: 5)\n", findBestValue([]int{2, 3, 5}, 10))
	fmt.Printf("%d (expected: 11361)\n", findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803))
}
```
