# 2553 — Separate The Digits In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SeparateTheDigitsInAnArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2553: Separate the Digits in an Array
// https://leetcode.com/problems/separate-the-digits-in-an-array/
// Difficulty: Easy
// Time O(n log m) | Space O(n log m)

import "fmt"

func main() {
	fmt.Println(SeparateTheDigitsInAnArray([]int{13, 25, 83, 77})) // [1,3,2,5,8,3,7,7]
	fmt.Println(SeparateTheDigitsInAnArray([]int{7, 1, 3, 9}))     // [7,1,3,9]
}

func SeparateTheDigitsInAnArray(nums []int) []int {
	res := []int{}
	for _, n := range nums {
		digits := []int{}
		for n > 0 {
			digits = append(digits, n%10)
			n /= 10
		}
		for i := len(digits) - 1; i >= 0; i-- {
			res = append(res, digits[i])
		}
	}
	return res
}
```
