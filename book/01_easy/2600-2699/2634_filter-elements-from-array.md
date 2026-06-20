# 2634 — Filter Elements From Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FilterElementsFromArray(arr []int, fn func(int, int) bool) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2634: Filter Elements from Array
// https://leetcode.com/problems/filter-elements-from-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Filters slice using a predicate.

import "fmt"

func main() {
	nums := []int{0, 10, 20, 30}
	greaterThan10 := func(n int, i int) bool { return n > 10 }
	fmt.Println(FilterElementsFromArray(nums, greaterThan10))

	nums2 := []int{1, 2, 3}
	firstIndex := func(n int, i int) bool { return i == 0 }
	fmt.Println(FilterElementsFromArray(nums2, firstIndex))
}

func FilterElementsFromArray(arr []int, fn func(int, int) bool) []int {
	result := []int{}
	for i, v := range arr {
		if fn(v, i) {
			result = append(result, v)
		}
	}
	return result
}
```
