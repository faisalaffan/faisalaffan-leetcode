# 0217 — Contains Duplicate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ContainsDuplicate(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsDuplicate(nums []int) bool {
  // HashMap: O(1) lookup
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
```
