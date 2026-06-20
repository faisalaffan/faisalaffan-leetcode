# 1979 — Find Greatest Common Divisor Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindGreatestCommonDivisorOfArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1979: Find Greatest Common Divisor of Array
// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{2, 5, 6, 9, 10})) // 2
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{7, 5, 6, 8, 3}))  // 1
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{3, 3}))            // 3
}

// Time: O(n), Space: O(1)
func FindGreatestCommonDivisorOfArray(nums []int) int {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	a, b := min, max
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
