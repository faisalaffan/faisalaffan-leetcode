# 3731 — Find Missing Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindMissingElements(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3731: Find Missing Elements
// https://leetcode.com/problems/find-missing-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMissingElements([]int{1, 4, 2, 5}))
	fmt.Println(FindMissingElements([]int{7, 8, 6, 9}))
	fmt.Println(FindMissingElements([]int{5, 1}))
}

// Time: O(n)
// Space: O(n)
func FindMissingElements(nums []int) []int {
  // HashMap: O(1) lookup
	has := make(map[int]bool)
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		has[v] = true
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}

  // Alokasi slice
	ans := make([]int, 0)
	for x := mn + 1; x < mx; x++ {
		if !has[x] {
			ans = append(ans, x)
		}
	}
	return ans
}
```
