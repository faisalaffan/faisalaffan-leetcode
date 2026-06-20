# 1608 — Special Array With X Elements Greater Than Or Equal X

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SpecialArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(log n) (for sorting)  
**Kompleksitas Ruang:** O(log n) (for sorting)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1608: Special Array With X Elements Greater Than or Equal X
// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func SpecialArray(nums []int) int {
  // Custom sort dengan comparator
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if nums[i] >= i+1 {
			if i == len(nums)-1 || nums[i+1] < i+1 {
				return i + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(SpecialArray([]int{3, 5}))
	fmt.Println(SpecialArray([]int{0, 0}))
	fmt.Println(SpecialArray([]int{0, 4, 3, 0, 4}))
}
```
