# 2226 — Maximum Candies Allocated To K Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumCandies(candies []int, k int64) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2226: Maximum Candies Allocated to K Children
// https://leetcode.com/problems/maximum-candies-allocated-to-k-children/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func maximumCandies(candies []int, k int64) int {
	lo, hi := 1, 0
	for _, c := range candies {
		if c > hi {
			hi = c
		}
	}

	result := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		var count int64 = 0
		for _, c := range candies {
			count += int64(c / mid)
		}
		if count >= k {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumCandies([]int{5, 8, 6}, 3))
	// Expected: 5

	// Test case 2
	fmt.Println(maximumCandies([]int{2, 5}, 11))
	// Expected: 0

	// Test case 3
	fmt.Println(maximumCandies([]int{4, 7, 5}, 16))
	// Expected: 1
}
```
