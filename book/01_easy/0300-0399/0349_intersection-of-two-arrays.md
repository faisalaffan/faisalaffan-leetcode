# 0349 — Intersection Of Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func IntersectionOfTwoArrays(nums1, nums2 []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n+m), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #349: Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n)
func IntersectionOfTwoArrays(nums1, nums2 []int) []int {
  // HashMap: O(1) lookup
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	var result []int
	for _, v := range nums2 {
		if set[v] {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
```
