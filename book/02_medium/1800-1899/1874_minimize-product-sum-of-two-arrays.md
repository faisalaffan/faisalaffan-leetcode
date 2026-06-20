# 1874 — Minimize Product Sum Of Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func MinProductSum(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1874: Minimize Product Sum of Two Arrays
// https://leetcode.com/problems/minimize-product-sum-of-two-arrays/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinProductSum([]int{5, 3, 4, 2}, []int{4, 2, 2, 5}))
	fmt.Println(MinProductSum([]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}))
}

// Time: O(n log n), Space: O(1)
func MinProductSum(nums1 []int, nums2 []int) int {
  // Sort O(n log n)
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	sum := 0
  // Linear scan O(n)
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i] * nums2[i]
	}
	return sum
}
```
