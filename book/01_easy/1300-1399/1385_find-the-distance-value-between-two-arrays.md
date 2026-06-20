# 1385 — Find The Distance Value Between Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findTheDistanceValue(arr1 []int, arr2 []int, d int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log m + m log m), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1385: Find the Distance Value Between Two Arrays
// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
// Difficulty: Easy
//
// LeetCode submission: func findTheDistanceValue(arr1 []int, arr2 []int, d int) int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTheDistanceValueBetweenTwoArrays([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)) // 2
	fmt.Println(FindTheDistanceValueBetweenTwoArrays([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3)) // 2
}

// Time: O(n log m + m log m), Space: O(1)
func FindTheDistanceValueBetweenTwoArrays(arr1 []int, arr2 []int, d int) int {
  // Sort O(n log n)
	sort.Ints(arr2)
	count := 0
	for _, v := range arr1 {
		if isFar(v, arr2, d) {
			count++
		}
	}
	return count
}

func isFar(v int, arr []int, d int) bool {
	idx := sort.SearchInts(arr, v)
	if idx < len(arr) && abs(arr[idx]-v) <= d {
		return false
	}
	if idx > 0 && abs(arr[idx-1]-v) <= d {
		return false
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
