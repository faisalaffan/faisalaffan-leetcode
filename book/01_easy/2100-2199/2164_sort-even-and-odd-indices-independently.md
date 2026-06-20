# 2164 — Sort Even And Odd Indices Independently

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SortEvenAndOddIndicesIndependently(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2164: Sort Even and Odd Indices Independently
// https://leetcode.com/problems/sort-even-and-odd-indices-independently/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortEvenAndOddIndicesIndependently([]int{4, 1, 2, 3})) // [2 3 4 1]
	fmt.Println(SortEvenAndOddIndicesIndependently([]int{2, 1}))       // [2 1]
}

// Time: O(n log n), Space: O(n)
func SortEvenAndOddIndicesIndependently(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	even := make([]int, 0, (n+1)/2)
  // Alokasi slice
	odd := make([]int, 0, n/2)

	for i, v := range nums {
		if i%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

  // Sort O(n log n)
	sort.Ints(even)
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

  // Alokasi slice
	result := make([]int, n)
	ei, oi := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = even[ei]
			ei++
		} else {
			result[i] = odd[oi]
			oi++
		}
	}
	return result
}
```
