# 2635 — Apply Transform Over Each Element In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ApplyTransformOverEachElementInArray(arr []int, fn func(int, int) int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2635: Apply Transform Over Each Element in Array
// https://leetcode.com/problems/apply-transform-over-each-element-in-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Maps a function over a slice.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	double := func(n int, i int) int { return n * 2 }
	fmt.Println(ApplyTransformOverEachElementInArray(nums, double))

	nums2 := []int{1, 2, 3}
	timesIndex := func(n int, i int) int { return n * i }
	fmt.Println(ApplyTransformOverEachElementInArray(nums2, timesIndex))
}

func ApplyTransformOverEachElementInArray(arr []int, fn func(int, int) int) []int {
  // Alokasi slice
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v, i)
	}
	return result
}
```
