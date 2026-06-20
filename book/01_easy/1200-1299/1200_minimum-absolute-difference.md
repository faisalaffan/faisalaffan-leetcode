# 1200 — Minimum Absolute Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumAbsDifference(arr []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1200: Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))       // [[1,2],[2,3],[3,4]]
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))  // [[1,3]]
}

// LeetCode submission: minimumAbsDifference
func minimumAbsDifference(arr []int) [][]int {
  // Sort O(n log n)
	sort.Ints(arr)
	minDiff := 1 << 31
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	var ans [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}
```
