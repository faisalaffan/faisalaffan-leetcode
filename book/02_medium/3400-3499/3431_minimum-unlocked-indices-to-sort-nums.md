# 3431 — Minimum Unlocked Indices To Sort Nums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func minUnlockedIndices(nums []int, locked []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3431: Minimum Unlocked Indices to Sort Nums
// https://leetcode.com/problems/minimum-unlocked-indices-to-sort-nums/
// Difficulty: Medium [Paid]
// Time: O(n^2) Space: O(n)

import (
	"fmt"
	"sort"
)

func minUnlockedIndices(nums []int, locked []int) int {
	n := len(nums)
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, nums)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)

	unlockCost := 0
	for i := 0; i < n; i++ {
		if nums[i] != sorted[i] && locked[i] == 1 {
			unlockCost++
		}
	}
	return unlockCost
}

func main() {
	fmt.Println(minUnlockedIndices([]int{1, 3, 2, 4}, []int{1, 1, 0, 1})) // 1
	fmt.Println(minUnlockedIndices([]int{1, 2, 3}, []int{1, 1, 1}))       // 0
}
```
