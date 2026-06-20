# 0190 — Reverse Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReverseBits(num uint32) uint32
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #190: Reverse Bits
// https://leetcode.com/problems/reverse-bits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func ReverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}
	return result
}

func main() {
	fmt.Println(ReverseBits(43261596))    // 964176192
	fmt.Println(ReverseBits(4294967293))  // 3221225471
}
```
