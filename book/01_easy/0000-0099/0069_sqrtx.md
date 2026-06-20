# 0069 — Sqrtx

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MySqrt(x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #69: Sqrt(x)
// https://leetcode.com/problems/sqrtx/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	lo, hi := 1, x/2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(MySqrt(4))
	fmt.Println(MySqrt(8))
	fmt.Println(MySqrt(0))
}
```
