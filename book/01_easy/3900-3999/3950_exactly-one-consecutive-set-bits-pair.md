# 3950 — Exactly One Consecutive Set Bits Pair

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ExactlyOneConsecutiveSetBitsPair(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3950: Exactly One Consecutive Set Bits Pair
// https://leetcode.com/problems/exactly-one-consecutive-set-bits-pair/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(6))
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(5))
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(3))
}

// Time: O(1)
// Space: O(1)
func ExactlyOneConsecutiveSetBitsPair(n int) bool {
	m := n & (n >> 1)
	return m > 0 && m&(m-1) == 0
}
```
