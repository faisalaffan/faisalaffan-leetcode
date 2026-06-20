# 0324 — Wiggle Sort Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func wiggleSort(nums []int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #324: Wiggle Sort II
// https://leetcode.com/problems/wiggle-sort-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	n := len(nums)
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, nums)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)

	// Fill from end of sorted array into odd positions first, then even
	mid := (n + 1) / 2
	j, k := mid-1, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			nums[i] = sorted[j]
			j--
		} else {
			nums[i] = sorted[k]
			k--
		}
	}
}

func main() {
	// Test case 1
	nums1 := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println("Test 1:", nums1)

	// Test case 2
	nums2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums2)
	fmt.Println("Test 2:", nums2)

	// Test case 3
	nums3 := []int{1, 2, 3}
	wiggleSort(nums3)
	fmt.Println("Test 3:", nums3)
}
```
