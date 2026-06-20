# 3701 — Compute Alternating Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ComputeAlternatingSum(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3701: Compute Alternating Sum
// https://leetcode.com/problems/compute-alternating-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ComputeAlternatingSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(ComputeAlternatingSum([]int{10, 5, 3}))
}

// Time: O(n)
// Space: O(1)
func ComputeAlternatingSum(nums []int) int {
	ans := 0
	for i, x := range nums {
		if i%2 == 0 {
			ans += x
		} else {
			ans -= x
		}
	}
	return ans
}
```
