# 0525 — Contiguous Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindMaxLength(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #525: Contiguous Array
// https://leetcode.com/problems/contiguous-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaxLength([]int{0, 1}))
	fmt.Println(FindMaxLength([]int{0, 1, 0}))
}

func FindMaxLength(nums []int) int {
	// Map count -> first index. count = (#1 - #0)
  // HashMap: O(1) lookup
	countMap := make(map[int]int)
	countMap[0] = -1
	count := 0
	maxLen := 0

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if prevIdx, ok := countMap[count]; ok {
			if i-prevIdx > maxLen {
				maxLen = i - prevIdx
			}
		} else {
			countMap[count] = i
		}
	}

	return maxLen
}
```
