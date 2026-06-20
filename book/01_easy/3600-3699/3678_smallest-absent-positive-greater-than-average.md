# 3678 — Smallest Absent Positive Greater Than Average

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SmallestAbsentPositiveGreaterThanAverage(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m) where m is the range of candidate values  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3678: Smallest Absent Positive Greater Than Average
// https://leetcode.com/problems/smallest-absent-positive-greater-than-average/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{3, 5}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{-1, 1, 2}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{4, -1}))
}

// Time: O(n + m) where m is the range of candidate values
// Space: O(n)
func SmallestAbsentPositiveGreaterThanAverage(nums []int) int {
  // HashMap: O(1) lookup
	has := make(map[int]bool)
	sum := 0
	for _, x := range nums {
		has[x] = true
		sum += x
	}

	avg := float64(sum) / float64(len(nums))
	ans := 1
	if int(avg)+1 > ans {
		ans = int(avg) + 1
	}

	for has[ans] {
		ans++
	}
	return ans
}
```
