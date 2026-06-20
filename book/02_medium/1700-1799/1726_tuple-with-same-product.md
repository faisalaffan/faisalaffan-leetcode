# 1726 — Tuple With Same Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func tupleSameProduct(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2), Space: O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1726: Tuple with Same Product
// https://leetcode.com/problems/tuple-with-same-product/
// Difficulty: Medium
// Time: O(n^2), Space: O(n^2)

import "fmt"

func tupleSameProduct(nums []int) int {
	n := len(nums)
  // HashMap: O(1) lookup
	productCount := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			productCount[product]++
		}
	}

	result := 0
	for _, count := range productCount {
		if count > 1 {
			// Each pair of pairs = 8 tuples (4! / 3 = 8)
			result += count * (count - 1) / 2 * 8
		}
	}
	return result
}

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))     // Expected: 8
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10})) // Expected: 16
	fmt.Println(tupleSameProduct([]int{1, 2, 3, 4, 6, 12})) // Expected: 40
}
```
