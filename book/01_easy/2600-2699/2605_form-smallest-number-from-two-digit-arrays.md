# 2605 — Form Smallest Number From Two Digit Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FormSmallestNumberFromTwoDigitArrays(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2605: Form Smallest Number From Two Digit Arrays
// https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FormSmallestNumberFromTwoDigitArrays([]int{4, 1, 3}, []int{5, 7}))       // 15
	fmt.Println(FormSmallestNumberFromTwoDigitArrays([]int{3, 5, 2, 6}, []int{3, 1, 7})) // 3
}

func FormSmallestNumberFromTwoDigitArrays(nums1 []int, nums2 []int) int {
	seen := [10]bool{}
	for _, n := range nums1 {
		seen[n] = true
	}

	common := 10
	for _, n := range nums2 {
		if seen[n] && n < common {
			common = n
		}
	}
	if common < 10 {
		return common
	}

	min1, min2 := 10, 10
	for _, n := range nums1 {
		if n < min1 {
			min1 = n
		}
	}
	for _, n := range nums2 {
		if n < min2 {
			min2 = n
		}
	}
	if min1 < min2 {
		return min1*10 + min2
	}
	return min2*10 + min1
}
```
