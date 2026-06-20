# 2919 — Minimum Increment Operations To Make Array Beautiful

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minIncrementOperations(nums []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2919: Minimum Increment Operations to Make Array Beautiful
// https://leetcode.com/problems/minimum-increment-operations-to-make-array-beautiful/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minIncrementOperations([]int{2, 3, 0, 0, 2}, 4))
	fmt.Println(minIncrementOperations([]int{0, 1, 3, 3}, 5))
	fmt.Println(minIncrementOperations([]int{1, 1, 2}, 1))
}

func minIncrementOperations(nums []int, k int) int64 {
	f, g, h := int64(0), int64(0), int64(0)
	for _, x := range nums {
		add := int64(0)
		if x < k {
			add = int64(k - x)
		}
		f, g, h = g, h, min64(f, min64(g, h))+add
	}
	return min64(f, min64(g, h))
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
```
