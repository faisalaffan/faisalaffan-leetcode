# 1913 — Maximum Product Difference Between Two Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumProductDifferenceBetweenTwoPairs(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1913: Maximum Product Difference Between Two Pairs
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{5, 6, 2, 7, 4}))      // 34
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{4, 2, 5, 9, 7, 4, 8})) // 64
}

// Time: O(n log n), Space: O(1) ignoring sort
func MaximumProductDifferenceBetweenTwoPairs(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
```
