# 3718 — Smallest Missing Multiple Of K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SmallestMissingMultipleOfK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + max_missing/k)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3718: Smallest Missing Multiple of K
// https://leetcode.com/problems/smallest-missing-multiple-of-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestMissingMultipleOfK([]int{8, 2, 3, 4, 6}, 2))
	fmt.Println(SmallestMissingMultipleOfK([]int{1, 4, 7, 10, 15}, 5))
}

// Time: O(n + max_missing/k)
// Space: O(n)
func SmallestMissingMultipleOfK(nums []int, k int) int {
  // HashMap: O(1) lookup
	has := make(map[int]bool)
	for _, v := range nums {
		has[v] = true
	}

	for x := k; ; x += k {
		if !has[x] {
			return x
		}
	}
}
```
