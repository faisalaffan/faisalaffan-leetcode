# 0561 — Array Partition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ArrayPartition(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #561: Array Partition
// https://leetcode.com/problems/array-partition/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func ArrayPartition(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	sum := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(ArrayPartition([]int{1, 4, 3, 2}))
	fmt.Println(ArrayPartition([]int{6, 2, 6, 5, 1, 2}))
}
```
