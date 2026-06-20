# 0035 — Search Insert Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SearchInsert(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #35: Search Insert Position
// https://leetcode.com/problems/search-insert-position/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func SearchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 7))
}
```
