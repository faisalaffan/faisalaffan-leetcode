# 1874 — Minimize Product Sum Of Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinProductSum(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	sum := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i] * nums2[i]
	}
	return sum
}
```
