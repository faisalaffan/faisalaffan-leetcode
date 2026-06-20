# 1356 — Sort Integers By The Number Of 1 Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func sortByBits(arr []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1356: Sort Integers by The Number of 1 Bits
// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
// Difficulty: Easy
//
// LeetCode submission: func sortByBits(arr []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortIntegersByTheNumberOfOneBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})) // [0 1 2 4 8 3 5 6 7]
	fmt.Println(SortIntegersByTheNumberOfOneBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1})) // [1 2 4 8 16 32 64 128 256 512 1024]
}

// Time: O(n log n), Space: O(1)
func SortIntegersByTheNumberOfOneBits(arr []int) []int {
  // Custom sort
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := popcount(arr[i]), popcount(arr[j])
		if bi != bj {
			return bi < bj
		}
		return arr[i] < arr[j]
	})
	return arr
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}
```
