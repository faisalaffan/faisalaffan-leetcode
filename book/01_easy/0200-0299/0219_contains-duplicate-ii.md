# 0219 — Contains Duplicate Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ContainsNearbyDuplicate(nums []int, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #219: Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsNearbyDuplicate(nums []int, k int) bool {
  // HashMap: O(1) lookup
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[n]; ok && i-j <= k {
			return true
		}
		seen[n] = i
	}
	return false
}

func main() {
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
```
