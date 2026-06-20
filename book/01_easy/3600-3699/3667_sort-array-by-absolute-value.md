# 3667 — Sort Array By Absolute Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SortArrayByAbsoluteValue(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3667: Sort Array By Absolute Value
// https://leetcode.com/problems/sort-array-by-absolute-value/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortArrayByAbsoluteValue([]int{3, -1, -4, 1, 5}))
}

// Time: O(n log n)
// Space: O(n)
func SortArrayByAbsoluteValue(nums []int) []int {
  // Custom sort dengan comparator
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})
	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
