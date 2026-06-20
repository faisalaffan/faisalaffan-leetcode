# 1064 — Fixed Point

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func fixedPoint(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1064: Fixed Point
// https://leetcode.com/problems/fixed-point/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))  // 3
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))    // 0
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9})) // -1
}

// LeetCode submission: fixedPoint
func fixedPoint(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] >= mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if arr[lo] == lo {
		return lo
	}
	return -1
}
```
