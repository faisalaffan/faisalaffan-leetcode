# 2387 — Median Of A Row Wise Sorted Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func matrixMedian(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** O(rows * log(cols) * log(max-min))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2387: Median of a Row Wise Sorted Matrix
// https://leetcode.com/problems/median-of-a-row-wise-sorted-matrix/
// Difficulty: Medium
// Time: O(rows * log(cols) * log(max-min)) | Space: O(1)
// Binary search on value, count elements <= mid.

import "fmt"

func main() {
	fmt.Println(matrixMedian([][]int{{1, 1, 2}, {2, 3, 3}, {1, 3, 4}})) // 2
	fmt.Println(matrixMedian([][]int{{1, 2}, {3, 4}}))                 // 2
}

func matrixMedian(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	target := r*c/2 + 1
	lo, hi := 1, 1000000

	for lo < hi {
		mid := (lo + hi) / 2
		count := 0
		for i := 0; i < r; i++ {
			// binary search in each row for count of elements <= mid
			row := grid[i]
			left, right := 0, c
  // Two-pointer: gerakkan kiri atau kanan
			for left < right {
				m := (left + right) / 2
				if row[m] <= mid {
					left = m + 1
				} else {
					right = m
				}
			}
			count += left
		}
		if count >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```
