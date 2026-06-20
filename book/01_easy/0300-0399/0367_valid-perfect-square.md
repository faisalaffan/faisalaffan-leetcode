# 0367 — Valid Perfect Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidPerfectSquare(num int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #367: Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ValidPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq == num {
			return true
		} else if sq < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(ValidPerfectSquare(16))
	fmt.Println(ValidPerfectSquare(14))
	fmt.Println(ValidPerfectSquare(1))
}
```
