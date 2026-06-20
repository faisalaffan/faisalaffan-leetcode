# 0342 — Power Of Four

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PowerOfFour(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #342: Power of Four
// https://leetcode.com/problems/power-of-four/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func PowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && (n-1)%3 == 0
}

func main() {
	fmt.Println(PowerOfFour(16))
	fmt.Println(PowerOfFour(5))
	fmt.Println(PowerOfFour(1))
}
```
