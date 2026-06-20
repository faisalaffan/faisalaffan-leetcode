# 1785 — Minimum Elements To Add To Form A Given Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minElements(nums []int, limit int, goal int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1785: Minimum Elements to Add to Form a Given Sum
// https://leetcode.com/problems/minimum-elements-to-add-to-form-a-given-sum/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	diff := goal - sum
	if diff < 0 {
		diff = -diff
	}

	// Minimum elements = ceil(diff / limit)
	return (diff + limit - 1) / limit
}

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4)) // Expected: 2
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0)) // Expected: 1
	fmt.Println(minElements([]int{0}, 1, 1000000)) // Expected: 1000000
}
```
