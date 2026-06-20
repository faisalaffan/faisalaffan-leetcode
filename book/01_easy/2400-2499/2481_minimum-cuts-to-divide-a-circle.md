# 2481 — Minimum Cuts To Divide A Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCutsToDivideACircle(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2481: Minimum Cuts to Divide a Circle
// https://leetcode.com/problems/minimum-cuts-to-divide-a-circle/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumCutsToDivideACircle(4)) // 2
	fmt.Println(MinimumCutsToDivideACircle(3)) // 3
	fmt.Println(MinimumCutsToDivideACircle(1)) // 0
}

func MinimumCutsToDivideACircle(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}
```
