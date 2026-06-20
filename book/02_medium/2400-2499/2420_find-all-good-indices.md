# 2420 — Find All Good Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func goodIndices(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2420: Find All Good Indices
// https://leetcode.com/problems/find-all-good-indices/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Precompute prefix non-increasing and suffix non-decreasing lengths.

import "fmt"

func main() {
	fmt.Println(goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2)) // [4, 5]
	fmt.Println(goodIndices([]int{1, 2, 3, 4, 5, 6}, 2))    // []
}

func goodIndices(nums []int, k int) []int {
	n := len(nums)
  // Alokasi slice
	pref := make([]int, n) // longest non-increasing ending at i
  // Alokasi slice
	suf := make([]int, n)  // longest non-decreasing starting at i

	pref[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			pref[i] = pref[i-1] + 1
		} else {
			pref[i] = 1
		}
	}

	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			suf[i] = suf[i+1] + 1
		} else {
			suf[i] = 1
		}
	}

  // Alokasi slice
	ans := make([]int, 0)
	for i := k; i < n-k; i++ {
		if pref[i-1] >= k && suf[i+1] >= k {
			ans = append(ans, i)
		}
	}
	return ans
}
```
