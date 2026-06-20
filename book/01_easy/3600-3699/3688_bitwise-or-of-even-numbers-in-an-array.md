# 3688 — Bitwise Or Of Even Numbers In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BitwiseOrOfEvenNumbersInAnArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3688: Bitwise OR of Even Numbers in an Array
// https://leetcode.com/problems/bitwise-or-of-even-numbers-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{7, 9, 11}))
	fmt.Println(BitwiseOrOfEvenNumbersInAnArray([]int{1, 8, 16}))
}

// Time: O(n)
// Space: O(1)
func BitwiseOrOfEvenNumbersInAnArray(nums []int) int {
	res := 0
	for _, n := range nums {
		if n%2 == 0 {
			res |= n
		}
	}
	return res
}
```
