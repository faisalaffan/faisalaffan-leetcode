# 0976 — Largest Perimeter Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestPerimeter(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(log n).  
**Kompleksitas Ruang:** O(log n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #976: Largest Perimeter Triangle
// https://leetcode.com/problems/largest-perimeter-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))          // 5
	fmt.Println(largestPerimeter([]int{1, 2, 1}))          // 0
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))       // 8
}

// largestPerimeter finds the largest perimeter of a triangle from the given side lengths.
// Time: O(n log n). Space: O(log n).
func largestPerimeter(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
```
