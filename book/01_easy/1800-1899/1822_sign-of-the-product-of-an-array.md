# 1822 — Sign Of The Product Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArraySign(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1822: Sign of the Product of an Array
// https://leetcode.com/problems/sign-of-the-product-of-an-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ArraySign(nums []int) int {
	negCount := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			negCount++
		}
	}
	if negCount%2 == 0 {
		return 1
	}
	return -1
}

func main() {
	fmt.Println(ArraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(ArraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(ArraySign([]int{-1, 1, -1, 1, -1}))
}
```
